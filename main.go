package main

import (
	"fmt"
	"os"
	exporter "github.com/alecrajeev/loggly-exporter/exporter"
	"github.com/alecrajeev/loggly-exporter/http"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	mets map[string]*prometheus.Desc
)

func init() {
	mets = exporter.AddMetrics()
}

func main() {
	fmt.Printf("Loggly Go Exporter\n")
	subDomain := os.Args[1]
	token := os.Args[2]

	exp := exporter.Exporter{
		APIMetrics: mets,
		Subdomain: subDomain,
		Token: token,
	}

	http.NewServer(exp).Start()
}