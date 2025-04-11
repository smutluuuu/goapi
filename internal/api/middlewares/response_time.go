package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	fmt.Println("Response Time Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Response Time Middleware being returned...")

		fmt.Println("Sent Response from ResponseTime")
		start := time.Now()

		//Create a custom ResponseWriter to capture the status code
		wrapperWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		duration := time.Since(start)
		w.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrapperWriter, r)
		// Calculate the duration

		// Log the request details
		duration = time.Since(start)
		fmt.Printf("Method: %s, URL: %s, Status:%d, Duration: %v\n", r.Method, r.URL, wrapperWriter.status, duration.String())
		fmt.Println("Sent Response from Response Time Middleware")

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
