package quiet_internet

import (
	"fmt"
	"io"
	"net/http"
)

const Base = "https://sizu.me/api/v1"

type Client struct {
	base       string
	httpClient *http.Client
}

type Option func(client *Client)

func New(token string, opts ...Option) *Client {
	hc := &http.Client{
		Transport: NewAuthorizationTransport(nil, token),
	}
	c := &Client{
		base:       Base,
		httpClient: hc,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			res.Body.Close()
			return nil, err
		}
		res.Body.Close()
		return nil, fmt.Errorf("HTTP error: %d, response: %s", res.StatusCode, string(body))
	}

	return res, nil
}

func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.httpClient.Transport = NewUserAgentTransport(c.httpClient.Transport, userAgent)
	}
}
