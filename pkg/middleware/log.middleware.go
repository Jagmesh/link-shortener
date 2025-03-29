package middleware

import (
	"fmt"
	"link-shortener/pkg/logger"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		rw := &responseWriter{w, http.StatusOK}

		starTime := time.Now()

		log.Info(fmt.Sprintf("[MIDDLEWARE] Request  | Method: %-6s | Path: %s", r.Method, r.URL.Path))
		next.ServeHTTP(rw, r)

		log.Info(fmt.Sprintf("[MIDDLEWARE] Response | Method: %-6s | Path: %-20s | Status: %d | Time: %s", r.Method, r.URL.Path, rw.statusCode, time.Since(starTime)))
	})
}
