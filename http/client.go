package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	client     *http.Client
	timeout    time.Duration
	retryCount int
}

const (
	MaxIdleConns       int  = 100
	MaxIdleConnections int  = 100
	RequestTimeout     int  = 30
	SSL                bool = true
)

func Header() http.Header {
	return http.Header{}
}
func NewClient(timeout int) *Client {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	return &Client{client: client}
}

func CustomClient(client *http.Client) *Client {

	return &Client{client: client}
}

func (c *Client) Get(url string, headers http.Header) ([]byte, int, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 500, err
	}

	request.Header = headers

	return c.Do(request)
}

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
