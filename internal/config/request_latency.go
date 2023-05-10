package config

import "github.com/prometheus/client_golang/prometheus"

type RequestLatency struct {
	HistogramVec *prometheus.HistogramVec
}

func NewRequestLatency() *RequestLatency {
	histogramVec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Latencia de las solicitudes HTTP en segundos.",
		Buckets: []float64{0.1, 0.2, 0.5, 1, 2, 5},
	}, []string{"method", "status_code"})

	return &RequestLatency{
		HistogramVec: histogramVec,
	}
}

func (r *RequestLatency) Observe(method string, path string, durationSeconds float64) {
	r.HistogramVec.WithLabelValues(method, path).Observe(durationSeconds)
}
