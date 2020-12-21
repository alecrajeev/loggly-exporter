package exporter

import (
	"fmt"
	"encoding/json"
)


// gatherData - Collects data from the API and stores into struct
func (e *Exporter) gatherData() ([]*Datum, error) {
	data := []*Datum{}

	responses, err := asyncHTTPGets(e.Subdomain, e.Token)

	if err != nil {
		return data, err
	}

	for _, response := range responses {

		// uses json encoding to parse the response
		d := new(Datum)
		errUnmarshal := json.Unmarshal(response.body, &d)

		if errUnmarshal != nil {
			fmt.Printf("Unable to parse json. Error: %v\n", errUnmarshal)
		}

		data = append(data, d)

	}

	return data, nil
}
