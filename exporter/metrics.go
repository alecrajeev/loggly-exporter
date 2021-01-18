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
		[]string{"name"}, nil,
	)

	return APIMetrics
}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(wrapperData []*WrapperDatum, ch chan <- prometheus.Metric) error {

	// APIMetrics - range through the wrapperData slice
	for _, w := range wrapperData {
		x := w.CountDatum
		ch <- prometheus.MustNewConstMetric(e.APIMetrics["Count"], prometheus.GaugeValue, float64(x.Count), w.Name)
	}

	return nil
}