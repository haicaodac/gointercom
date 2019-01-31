package gointercom

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get ...
func Get(url string) ([]byte, error) {
	var cfg = GetConfig()

	var AccessToken = cfg.GetToken()
	var bearer = "Bearer " + AccessToken

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	// Send req using http Client
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return []byte(body), nil
}

// Post ...
func Post(url string, data []byte) ([]byte, error) {
	var cfg = GetConfig()

	var AccessToken = cfg.GetToken()
	var bearer = "Bearer " + AccessToken

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return body, nil
}
