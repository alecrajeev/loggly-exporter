package exporter

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
)

type SearchQuery struct {
	Name string
	Query string
}

type Exporter struct {
	APIMetrics 			map[string]*prometheus.Desc
	Subdomain 			string
	Token 	     		string
	ListenerPort		int
	Query 				string
	SearchQueries       []SearchQuery
}


// Datum is used to store data from around the number of loggly events
type Datum struct {
	Count int `json:"count"`
}

type WrapperDatum struct {
	CountDatum Datum
	Name 	   string
}

// Response struct is used to store http.Response and associated data
type Response struct {
	subdomain string
	name 	  string
	response  *http.Response
	body 	  []byte
	err 	  error

}