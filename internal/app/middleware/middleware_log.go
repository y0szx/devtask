package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Logger is a middleware that logs details about each HTTP request and its duration.
func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		const maxLogFileSize = 10 * 1024 * 1024 // 10 MB
		const logFileName = "server.log"

		// Function to get the size of the log file
		getFileSize := func(fileName string) int64 {
			info, err := os.Stat(fileName)
			if err != nil {
				return 0
			}
			return info.Size()
		}

		// Function to rotate the log file
		rotateLogFile := func(fileName string) error {
			timestamp := time.Now().Format("20060102-150405")
			newFileName := fmt.Sprintf("%s.%s", fileName, timestamp)
			return os.Rename(fileName, newFileName)
		}

		// Check the size of the current log file and rotate if necessary
		if getFileSize(logFileName) >= maxLogFileSize {
			if err := rotateLogFile(logFileName); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

		// Open or create the log file
		logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer logFile.Close()

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

		logEntry := fmt.Sprintf(
			"Method: %s\nHeader: %v\nPath: %s\nBody: %s\nTimeStamp: %s\n\n",
			req.Method, req.Header, req.URL.Path, string(bodyBytes), start,
		)

		// Write the log entry to the file
		if _, err := logFile.WriteString(logEntry); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, req)
	}
}
