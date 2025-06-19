package middleware

import (
	"log/slog"
	"net/http"
)

type HeaderWriter struct {
	logger *slog.Logger
}

func NewHeaderWriter(log *slog.Logger) *HeaderWriter {
	return &HeaderWriter{logger: log.With(slog.String("module", "middleware.Header"))}
}

func (hw *HeaderWriter) WriteHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hw.WriteContentTypeHeader(w.Header())
		next.ServeHTTP(w, r)
	})
}

func (hw *HeaderWriter) WriteContentTypeHeader(h http.Header) {
	h.Set("Content-Type", "application/json")
}
