package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"reflect"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"github.com/prometheus/client_golang/prometheus"
)

// Prometheus metrics
var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nt_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method"},
	)

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "nt_http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.ExponentialBuckets(0.001, 10, 5),
		},
		[]string{"path", "method"},
	)

	ActiveConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "nt_active_connections",
			Help: "Number of active connections",
		},
	)
)

type Instrumentation struct {
	logger *slog.Logger
}

func NewInstrumentation(logger *slog.Logger) *Instrumentation {
	i := &Instrumentation{}
	i.logger = logger.With("module", reflect.TypeOf(i))
	return i
}

func (i *Instrumentation) Instrument(next nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
		path := r.URL.Path
		method := r.Method
		timer := prometheus.NewTimer(HttpRequestDuration.WithLabelValues(path, method))
		HttpRequestsTotal.WithLabelValues(path, method).Inc()
		ActiveConnections.Inc()
		res, err := next(ctx, w, r, request)
		timer.ObserveDuration()
		ActiveConnections.Dec()
		return res, err
	}
}
