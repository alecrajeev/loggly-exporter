package http

import (
	exporter "github.com/alecrajeev/loggly-exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Server struct {
	Handler http.Handler
	exporter exporter.Exporter
}

func NewServer(exporter exporter.Exporter) *Server {
	r := http.NewServeMux()

	// Register metrics for the search query
	// This invokes the Collect method through the prometheus client libraries
	prometheus.MustRegister(&exporter)

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
		                <head><title>Loggly Exporter</title></head>
		                <body>
		                   <h1>Loggly Prometheus Metrics Exporter</h1>
						   <p>For more information, visit <a href=https://github.com/alecrajeev/loggly-exporter>Loggly</a></p>
		                   <p><a href='/metrics'>Metrics</a></p>
		                   </body>
		                </html>
		              `))
	})

	return &Server{Handler: r, exporter: exporter}
}

func (s *Server) Start() {
	http.ListenAndServe(":"+"9786", s.Handler)
}





