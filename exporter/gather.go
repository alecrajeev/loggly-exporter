package exporter

import (
	"fmt"
	"encoding/json"
)


// gatherData - Collects data from the API and stores into struct
func (e *Exporter) gatherData() ([]*WrapperDatum, error) {
	wrapperData := []*WrapperDatum{}

	responses, err := asyncHTTPGets(e.Subdomain, e.Token, e.SearchQueries)

	if err != nil {
		return wrapperData, err
	}

	for _, response := range responses {

		// uses json encoding to parse the response
		d := new(Datum)
		errUnmarshal := json.Unmarshal(response.body, &d)

		if errUnmarshal != nil {
			fmt.Printf("Unable to parse json. Error: %v\n", errUnmarshal)
		}

		wD := new(WrapperDatum)
		wD.Name = response.name
		wD.CountDatum = *d

		wrapperData = append(wrapperData, wD)
	}

	return wrapperData, nil
}
