package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"reflect"
	"time"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
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

func (rm *RequestMetadata) Time(next nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
		start := time.Now()
		writer := &StatusWriter{ResponseWriter: w, statusCode: http.StatusOK}
		res, err := next(ctx, w, r, request)
		rm.logger.Info(r.Method, slog.Int("status", writer.statusCode), slog.String("path", r.URL.Path), slog.Any("time", time.Since(start)))
		return res, err
	}
}
