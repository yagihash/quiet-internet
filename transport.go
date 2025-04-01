package quiet_internet

import (
	"fmt"
	"net/http"
)

type AuthorizationTransport struct {
	rt    http.RoundTripper
	token string
}

func NewAuthorizationTransport(parent http.RoundTripper, token string) *AuthorizationTransport {
	return &AuthorizationTransport{
		rt:    parent,
		token: token,
	}
}

func (t *AuthorizationTransport) parent() http.RoundTripper {
	if t.rt == nil {
		return http.DefaultTransport
	}

	return t.rt
}

func (t *AuthorizationTransport) CancelRequest(req *http.Request) {
	type canceler interface {
		CancelRequest(*http.Request)
	}

	if c, ok := t.parent().(canceler); ok {
		c.CancelRequest(req)
	}
}

func (t *AuthorizationTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", t.token))

	return t.parent().RoundTrip(req)
}

type UserAgentTransport struct {
	rt        http.RoundTripper
	userAgent string
}

func NewUserAgentTransport(parent http.RoundTripper, userAgent string) *UserAgentTransport {
	return &UserAgentTransport{
		rt:        parent,
		userAgent: userAgent,
	}
}

func (t *UserAgentTransport) parent() http.RoundTripper {
	if t.rt == nil {
		return http.DefaultTransport
	}

	return t.rt
}

func (t *UserAgentTransport) CancelRequest(req *http.Request) {
	type canceler interface {
		CancelRequest(*http.Request)
	}

	if c, ok := t.parent().(canceler); ok {
		c.CancelRequest(req)
	}
}

func (t *UserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("user-agent", t.userAgent)

	return t.parent().RoundTrip(req)
}
