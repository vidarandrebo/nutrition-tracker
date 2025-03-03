package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type RequestMetadata struct {
	log *slog.Logger
}

type StatusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (sw *StatusWriter) WriteHeader(statusCode int) {
	sw.ResponseWriter.WriteHeader(statusCode)
	sw.statusCode = statusCode
}
func NewRequestTimer(log *slog.Logger) *RequestMetadata {
	return &RequestMetadata{log: log.With(slog.String("module", "middleware.RequestMetadata"))}
}

func (rt *RequestMetadata) Time(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		writer := &StatusWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(writer, r)
		rt.log.Info(r.Method, slog.Int("status", writer.statusCode), slog.String("path", r.URL.Path), slog.Any("time", time.Since(start)))
	})
}
