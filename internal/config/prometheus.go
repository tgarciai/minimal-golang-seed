package config

import (
	"github.com/prometheus/client_golang/prometheus"
)

func SetupPrometheus(requestLatency *RequestLatency) *prometheus.Registry {

	registry := prometheus.NewRegistry()
	registry.MustRegister(prometheus.NewGoCollector())
	registry.MustRegister(requestLatency.HistogramVec)

	return registry
}
