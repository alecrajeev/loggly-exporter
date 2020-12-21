package main

import (
	"fmt"
	exporter "github.com/alecrajeev/loggly-exporter/exporter"
	"os"
)

func main() {
	fmt.Printf("Loggly Go Exporter\n")
	subDomain := os.Args[1]
	token := os.Args[2]


	exporter.Collect(subDomain, token)
}