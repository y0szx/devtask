package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return
		}
		req.Body = io.NopCloser(io.Reader(bytes.NewBuffer(bodyBytes)))
		fmt.Println("Method: ", req.Method)
		fmt.Println("Header: ", req.Header)
		fmt.Println("Path: ", req.URL.Path)
		fmt.Println("Body:", string(bodyBytes))
		fmt.Printf("TimeStamp: %s\n\n", start)
		next.ServeHTTP(w, req)
	}
}
