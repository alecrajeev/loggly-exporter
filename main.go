package main

import (
	"fmt"
	"flag"
	"strconv"
	exporter "github.com/alecrajeev/loggly-exporter/exporter"
	conf "github.com/alecrajeev/loggly-exporter/config"
	"github.com/alecrajeev/loggly-exporter/http"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	mets 				  	 map[string]*prometheus.Desc
	configFile            	 = flag.String("config.file", "config.yml", "Path to configuration file.")
	logglyToken 			 = flag.String("token", "", "Loggly API token")

	applicationConf          = conf.Conf{}
)

func init() {
	mets = exporter.AddMetrics()
}

func main() {

	flag.Parse()

	fmt.Println("Parsing config..")
	if err := applicationConf.Load(configFile); err != nil {
		fmt.Println("Failed to parse config file")
		return
	}

	logglySubdomain := applicationConf.LogglySubDomain

	fmt.Printf("Loggly Subdomain: %v\n", logglySubdomain)

	var logglySearches []exporter.SearchQuery
	for _, l := range applicationConf.LogglySearches {
		logglySearches = append(logglySearches, exporter.SearchQuery {
			Name: l.Name,
			Query: l.Query,
		})
	}
	if len(logglySearches) < 1 {
		fmt.Printf("Need to add loggly_searches\n")
		return
	}

	ListenerPort, err := strconv.Atoi(applicationConf.ListenerPort)
	if err != nil {
		fmt.Println("Got error parsing listener port")
		return
	}

	fmt.Printf("Loggly token: %v\n", *logglyToken)
	fmt.Printf("Listner Port: %v\n", ListenerPort)

	fmt.Printf("Loggly Go Exporter\n")

	exp := exporter.Exporter{
		APIMetrics: mets,
		Subdomain: logglySubdomain,
		Token: *logglyToken,
		ListenerPort: ListenerPort,
		SearchQueries: logglySearches,
	}

	http.NewServer(exp).Start()
}