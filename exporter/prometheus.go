package exporter

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)


// Describe - loops through API metrics and passes them to prometheus.Describe
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	for _,m := range e.APIMetrics {
		ch <- m
	}

}

// Collect - called on by the Prometheus Client library
// This function is called when a scrape is performed on the /metrics page
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	wrapperData := []*WrapperDatum{}
	var err error

	wrapperData, err = e.gatherData()
	if err != nil {
		fmt.Println("Error getting data")
		return
	}

	// Set prometheus gauge metrics using the data gathered
	errPrometheus := e.processMetrics(wrapperData, ch)

	if errPrometheus != nil {
		fmt.Println("Error processing metrics,", errPrometheus)
		return
	}

}





