package Newhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

const (
	MaxIdleConns       int  = 100
	MaxIdleConnections int  = 100
	RequestTimeout     int  = 30
	SSL                bool = true
)

// Header is a function that will response http.Header , this function will set header for http
func Header() http.Header {
	return http.Header{}
}

// NewClient is a function of default http client with add timeout
// This function will return Struct Client that containt http client
func NewClient(timeout int) *Client {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	return &Client{client: client}
}

// This function is custom http client , you can add transport , proxy or etc
// This function will return custom http client , this function will return custom http
func CustomClient(client *http.Client) *Client {

	return &Client{client: client}
}

// This function is custom http client with method get  , with parameter url and header
// This function will return :
// - data (result of endpoint that has  byte type that can unmarshal with the struct)
// - Status code (status code of response end point that has type int)
// - Error (Error response endpoint that has type error)
func (c *Client) Get(url string, headers http.Header) ([]byte, int, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 500, err
	}

	request.Header = headers

	return c.Do(request)
}

// This function is custom http client with method post  , with parameter url ,header and body payload
// This function will return :
// - data (result of endpoint that has  byte type that can unmarshal with the struct)
// - Status code (status code of response end point that has type int)
// - Error (Error response endpoint that has type error)
func (c *Client) Post(url string, headers http.Header, payload interface{}) ([]byte, int, error) {

	reqBody, err := json.Marshal(&payload)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, 500, err
	}
	defer request.Body.Close()
	request.Header = headers

	return c.Do(request)
}

// This function is custom http client with method put  , with parameter url ,header and body payload
// This function will return :
// - data (result of endpoint that has  byte type that can unmarshal with the struct)
// - Status code (status code of response end point that has type int)
// - Error (Error response endpoint that has type error)
func (c *Client) Put(url string, headers http.Header, payload interface{}) ([]byte, int, error) {

	reqBody, err := json.Marshal(&payload)
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, 500, err
	}
	defer request.Body.Close()
	request.Header = headers
	return c.Do(request)
}

// This function is custom http client with method put  , with parameter url ,header and body payload
// This function will return :
// - data (result of endpoint that has  byte type that can unmarshal with the struct)
// - Status code (status code of response end point that has type int)
// - Error (Error response endpoint that has type error)
func (c *Client) Patch(url string, headers http.Header, payload interface{}) ([]byte, int, error) {

	reqBody, err := json.Marshal(&payload)
	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, 500, err
	}
	defer request.Body.Close()
	request.Header = headers

	return c.Do(request)
}

// This function is custom http client with method put  , with parameter url ,header and body payload
// This function will return :
// - data (result of endpoint that has  byte type that can unmarshal with the struct)
// - Status code (status code of response end point that has type int)
// - Error (Error response endpoint that has type error)
func (c *Client) Delete(url string, headers http.Header) ([]byte, int, error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, 500, err
	}
	defer request.Body.Close()
	request.Header = headers

	return c.Do(request)
}

func (c *Client) Do(request *http.Request) ([]byte, int, error) {

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, 500, err
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, 500, err
	}

	io.Copy(ioutil.Discard, resp.Body)

	return body, resp.StatusCode, nil
}
