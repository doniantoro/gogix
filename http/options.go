package http

import "time"

type Option func(*Client)

func WithHTTPTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}
