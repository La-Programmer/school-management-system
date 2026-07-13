package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Create a custom response writer to capture status code
		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(wrappedWriter, r)

		// Calculate duration
		duration := time.Since(startTime)

		wrappedWriter.Header().Set("X-Response-Time", duration.String())

		// Log request details
		fmt.Printf("Method: %s, URL: %s, Status: %d, Duration: %v\n", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Sent Response from response time middleware")
	})
}

// response writer
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
