package exporter

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)


func Collect2(subDomain string, token string) {

	fmt.Printf("Collect metrics here\n")
	query := "json.level(\"INFO\")"

	err := getResponse(subDomain, query, token)
	_ = err

}


func getResponse(logglySubDomain string, query string, token string) error {

	resp, err := getHTTPResponse(logglySubDomain, query, token)

	if err != nil {
		fmt.Printf("Got error in response")
		return err
	}

	// closes response body when done reading it (prevents memory leaks)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when reading body")
		return err
	}

	var d Datum

	err2 := json.Unmarshal(body, &d)

	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	fmt.Println(d)


	return err

}

func getHTTPResponse(logglySubDomain string, query string, token string) (*http.Response, error) {


	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := "https://" + logglySubDomain + ".loggly.com/apiv2/events/count"
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
