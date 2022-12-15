package main

import (
	"fmt"

	"github.com/doniantoro/gogix/helper"
	http "github.com/doniantoro/gogix/http"
)

// go mod init github.com/doniantoro/gogix

const (
	baseURL = "http://localhost:3000/v3/digicore/check-coverage"
)

func main() {

	httpClient := http.NewClient(20)
	header := http.Header()
	header.Set("Content-Type", "application/json")
	response, code, err := httpClient.Get(baseURL, header)
	if err != nil {

	}
	if code == 200 {
		fmt.Println("Response: %s", string(response))

	} else {
		fmt.Println(helper.MapError(code))
	}
	// http.MapError("test")
	// helper.ResponseErrorWithJSON(helper.ErrBadGateway, "Masuk")
}
