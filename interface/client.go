package interfaces

import (
	"net/http"
)

type Client interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, body interface{}, headers http.Header) (*http.Response, error)
	Put(url string, body interface{}, headers http.Header) (*http.Response, error)
	Patch(url string, body interface{}, headers http.Header) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	Do(req *http.Request) (*http.Response, error)
}
