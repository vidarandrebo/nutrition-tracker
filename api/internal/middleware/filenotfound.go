package middleware

import (
	"log/slog"
	"net/http"
)

type FileNotFoundInterceptor struct {
	log *slog.Logger
}

type FileNotFoundResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *FileNotFoundResponseWriter) Write(b []byte) (int, error) {
	// prevent the default 404 not found being written to body
	if w.statusCode == http.StatusNotFound {
		return len(b), nil
	}
	if w.statusCode != 0 {
		w.WriteHeader(w.statusCode)
	}
	return w.ResponseWriter.Write(b)
}

func (w *FileNotFoundResponseWriter) WriteHeader(statusCode int) {
	if statusCode >= 300 && statusCode < 400 {
		w.ResponseWriter.WriteHeader(statusCode)
		return
	}
	w.statusCode = statusCode
}

func NewFileNotFoundInterceptor(log *slog.Logger) *FileNotFoundInterceptor {
	return &FileNotFoundInterceptor{log: log.With(slog.String("module", "middleware.FileNotFoundInterceptor"))}
}

func (fi *FileNotFoundInterceptor) RespondWithFallback(next http.Handler, path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &FileNotFoundResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(writer, r)
		if writer.statusCode == http.StatusNotFound {
			r.URL.Path = path
			w.Header().Set("Content-Type", "text/html")
			next.ServeHTTP(w, r)
		}
	})
}
