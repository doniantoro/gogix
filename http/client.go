package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
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
	// Header             http.Header = http.Header{}
)

func Header() http.Header {
	return http.Header{}
}
func NewClient(timeout int) *Client {

	timeoutConv, _ := time.ParseDuration(strconv.Itoa(123))
	client := &http.Client{Timeout: timeoutConv * time.Second}

	return &Client{client: client}
}
func (c *Client) Get(url string, headers http.Header) ([]byte, int, error) {

	fmt.Println(c.client)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 500, err
	}

	request.Header = headers

	return c.Do(request)
}

func (c *Client) Post(url string, headers http.Header, payload interface{}) ([]byte, int, error) {

	fmt.Println(c.client)
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

	fmt.Println(c.client)
	reqBody, err := json.Marshal(&payload)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, 500, err
	}
	defer request.Body.Close()
	request.Header = headers

	return c.Do(request)
}

func (c *Client) Do(request *http.Request) ([]byte, int, error) {

	resp, err := c.client.Do(request)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, 500, err
	}

	io.Copy(ioutil.Discard, resp.Body)
	// defer request.Body.Close()

	return body, resp.StatusCode, nil
}
