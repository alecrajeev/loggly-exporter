package exporter

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	APIMetrics map[string]*prometheus.Desc
	Subdomain string
	Token string
	ListenerPort int
	Query string
}


// Datum is used to store data from around the number of loggly events
type Datum struct {
	Count int `json:"count"`
}

// Response struct is used to store http.Response and associated data
type Response struct {
	subdomain string
	response *http.Response
	body []byte
	err error

}