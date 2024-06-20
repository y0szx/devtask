package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Logger is a middleware that logs details about each HTTP request and its duration.
func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Record the start time of the request
		start := time.Now()

		// Read the body of the request
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return
		}

		// Restore the body to the request after reading it
		req.Body = io.NopCloser(io.Reader(bytes.NewBuffer(bodyBytes)))

		// Log the HTTP method
		fmt.Println("Method: ", req.Method)

		// Log the request headers
		fmt.Println("Header: ", req.Header)

		// Log the request path
		fmt.Println("Path: ", req.URL.Path)

		// Log the request body
		fmt.Println("Body:", string(bodyBytes))

		// Log the timestamp of the request
		fmt.Printf("TimeStamp: %s\n\n", start)

		// Call the next handler in the chain
		next.ServeHTTP(w, req)
	}
}
