package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	APIMetrics map[string]*prometheus.Desc
}


// Datum is used to store data from around the number of loggly events
type Datum struct {
	Count int `json:"count"`
}