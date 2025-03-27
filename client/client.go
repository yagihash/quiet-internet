package client

import (
	"net/http"

	"github.com/yagihash/quiet-internet/client/transport"
)

const Base = "https://sizu.me/api/"

type Client struct {
	base       string
	httpClient *http.Client
}

type Option func(client *Client)

func New(token string, opts ...Option) *Client {
	hc := &http.Client{
		Transport: transport.NewAuthorizationTransport(nil, token),
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

func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.httpClient.Transport = transport.NewUserAgentTransport(c.httpClient.Transport, userAgent)
	}
}
