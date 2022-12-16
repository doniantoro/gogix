package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/doniantoro/gogix/NewHelper"
	gogix "github.com/doniantoro/gogix/NewHttp"
	"github.com/gorilla/mux"
)

const (
	// baseURL = "http://localhost:3000/v3/digicore/check-coverage"
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

	r := mux.NewRouter()
	r.HandleFunc("/success", FuncSuccess).Methods("GET")
	r.HandleFunc("/failed", FuncFailed).Methods("GET")
	r.HandleFunc("/failed-message", FuncFailedWithCustomeMessage).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8010", r)

}

func ExampleGet() {
	httpClient := gogix.NewClient(20)
	header := gogix.Header()
	header.Set("Content-Type", "application/json")
	response, code, err := httpClient.Get(baseURL, header)
	if err != nil {
		fmt.Println(NewHelper.MapError(code))
	}
	if code == 200 {
		fmt.Println("Response: %s", string(response))
	} else {
		fmt.Println(NewHelper.MapError(code))
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

		fmt.Println(NewHelper.MapError(code))
	}
	if code == 200 {
		fmt.Println("Response: %s", string(response))

	} else {
		fmt.Println(NewHelper.MapError(code))
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

		fmt.Println(NewHelper.MapError(code))
	}
	if code == 200 {
		fmt.Println("Response: %s", string(response))

	} else {
		fmt.Println(NewHelper.MapError(code))
	}
}

func FuncSuccess(w http.ResponseWriter, r *http.Request) {
	order := Order{}
	order.Msisdn = "089526265660"
	NewHelper.ResponseWithJSON(w, 200, "Success Get Data", order)
}

func FuncFailed(w http.ResponseWriter, r *http.Request) {

	NewHelper.ResponseErrorWithJSON(w, NewHelper.ErrBadRequest, "")
}

func FuncFailedWithCustomeMessage(w http.ResponseWriter, r *http.Request) {

	NewHelper.ResponseErrorWithJSON(w, NewHelper.ErrBadRequest, "Error With Custome Message")
}
