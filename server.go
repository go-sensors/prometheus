package prometheus

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsServer is an HTTP serve for serving Prometheus metrics
type MetricsServer struct {
	addr string
}

// NewMetricsServer returns a MetricsServer
func NewMetricsServer(addr string) *MetricsServer {
	return &MetricsServer{
		addr,
	}
}

// Run starts the HTTP server and handles requests to /metrics
func (s *MetricsServer) Run(ctx context.Context) error {
	handler := http.NewServeMux()
	handler.Handle("/metrics", promhttp.Handler())

	metricServer := http.Server{
		Addr:    s.addr,
		Handler: handler,
	}

	go func() {
		<-ctx.Done()
		metricServer.Close()
	}()

	err := metricServer.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}
