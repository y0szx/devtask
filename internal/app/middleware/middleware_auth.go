package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

// BasicAuth is a middleware function that enforces basic authentication using a provided username and password.
func BasicAuth(next http.Handler, expectedUsername string, expectedPassword string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Retrieve the username and password from the request's Basic Authentication header
		username, password, ok := req.BasicAuth()

		if ok {
			// Hash the provided username and password
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))

			// Hash the expected username and password
			expectedUsernameHash := sha256.Sum256([]byte(expectedUsername))
			expectedPasswordHash := sha256.Sum256([]byte(expectedPassword))

			// Compare the hashed provided credentials with the hashed expected credentials
			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			// If both the username and password match, proceed to the next handler
			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, req)
			} else {
				// If the credentials do not match, respond with an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			// If the Basic Authentication header is missing or invalid, respond with an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
