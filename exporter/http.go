package exporter

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

func asyncHTTPGets(subDomain string, token string, query string) ([]*Response, error) {

	queryCount := 1

	// Channels used to make concurrent requests
	ch := make(chan *Response, queryCount)

	responses := []*Response{}

	go func(subDomain string) {
		err := getResponse(subDomain, token, query, ch)

		if err != nil {
			ch <- &Response{subDomain, nil, []byte{}, err}
		}
	}(subDomain)

	for {
		select {
		case r := <-ch:
			if r.err != nil {
				fmt.Printf("Error scraping API, Error: %v\n", r.err)
				break
			}
			responses = append(responses, r)

			if len(responses) == queryCount {
				return responses, nil
			}
		}
	}
}

// getResponse collects an individual http.response and returns a *Response
func getResponse(subDomain string, token string, query string, ch chan<- *Response) error {

	resp, err := getHTTPResponse(subDomain, token, query)

	if err != nil {
		fmt.Println("Error getting response", err)
		return err
	}

	// closes response body when done reading it (prevents memory leaks)
	defer resp.Body.Close()

	body, errIo := ioutil.ReadAll(resp.Body)
	if errIo != nil {
		fmt.Println("Error when reading body")
		return errIo
	}

	ch <- &Response{subDomain, resp, body, errIo}

	return nil
}

// getHTTPResponse handles the http client creation, token setting and returns the *http.response
func getHTTPResponse(subDomain string, token string, query string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := "https://" + subDomain + ".loggly.com/apiv2/events/count"
	url = url + "?q=" + query

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("got error from http")
		return nil, err
	} else {
		fmt.Println(req)
	}

	header_token := "Bearer " + token

	req.Header.Set("Authorization", header_token)

	resp, err := client.Do(req)

	return resp, err
}