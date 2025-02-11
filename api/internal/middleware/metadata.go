package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type RequestMetadata struct {
	log *slog.Logger
}

func NewRequestTimer(log *slog.Logger) *RequestMetadata {
	return &RequestMetadata{log: log.With(slog.String("module", "middleware.RequestMetadata"))}
}

func (rt *RequestMetadata) Time(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		rt.log.Info(r.Method, slog.String("path", r.URL.Path), slog.Any("time", time.Since(start)))
	})
}
