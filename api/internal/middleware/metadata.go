package middleware

import (
	"log/slog"
	"net/http"
	"reflect"
	"time"
)

type RequestMetadata struct {
	logger *slog.Logger
}

type StatusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (sw *StatusWriter) WriteHeader(statusCode int) {
	sw.ResponseWriter.WriteHeader(statusCode)
	sw.statusCode = statusCode
}

func NewRequestMetadata(logger *slog.Logger) *RequestMetadata {
	rm := &RequestMetadata{}
	rm.logger = logger.With("module", reflect.TypeOf(rm))
	return rm
}

func (rt *RequestMetadata) Time(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		writer := &StatusWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(writer, r)
		rt.logger.Info(r.Method, slog.Int("status", writer.statusCode), slog.String("path", r.URL.Path), slog.Any("time", time.Since(start)))
	})
}
