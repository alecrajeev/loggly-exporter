package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

// AddMetrics - Adds all of the metrics to a map of strings, then returns the map
func AddMetrics() map[string]*prometheus.Desc {
	APIMetrics := make(map[string]*prometheus.Desc)

	APIMetrics["Count"] = prometheus.NewDesc(
		prometheus.BuildFQName("loggly", "search", "count"),
		"Number events in loggly for a particular search",
		[]string{}, nil,
	)

	return APIMetrics
}