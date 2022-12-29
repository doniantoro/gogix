package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/doniantoro/gogix"
)

const (
	baseURL = "http://localhost:3013/order/omni/credit"
)

type Order struct {
	Msisdn string
}

func main() {
	ExampleGet()
	// ExamplePost()
	// ExampleCustomClient()
	// ExampleCustomClient()

}

func ExampleGet() {
	httpClient := gogix.NewClient(20)
	header := gogix.Header()
	header.Set("Content-Type", "application/json")
	response, code, err := httpClient.Get(baseURL, header)
	if err != nil {
		fmt.Println(err)
	}
	if code == 200 {
		fmt.Println(string(response))
	} else {
		fmt.Println(code)
	}
}

func ExamplePost() {
	httpClient := gogix.NewClient(20)
	header := gogix.Header()
	header.Set("Content-Type", "application/json")
	header.Set("AccessToken", "test")
	payload := Order{}
	payload.Msisdn = "089526265660"

	response, code, err := httpClient.Post(baseURL, header, payload)
	if err != nil {
		fmt.Println(err)
	}
	if code == 200 {
		fmt.Println(string(response))
	} else {
		fmt.Println(code)
	}
}

func ExampleCustomClient() {

	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(30) * time.Second,
	}

	httpClient := gogix.CustomClient(client)
	header := gogix.Header()
	header.Set("Content-Type", "application/json")
	response, code, err := httpClient.Get(baseURL, header)
	if err != nil {
		fmt.Println(err)
	}
	if code == 200 {
		fmt.Println(string(response))
	} else {
		fmt.Println(code)
	}
}
