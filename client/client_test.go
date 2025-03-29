package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		testToken := "test-token"
		want := &Client{
			base: Base,
		}
		got := New(testToken)

		if !cmp.Equal(want, got, cmp.Comparer(func(a, b *Client) bool {
			return a.base == b.base
		})) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("WithOption", func(t *testing.T) {
		testToken := "test-token"
		want := &Client{
			base: "test-base",
		}
		got := New(testToken, func(client *Client) {
			client.base = "test-base"
		})
		if !cmp.Equal(want, got, cmp.Comparer(func(a, b *Client) bool {
			return a.base == b.base
		})) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestWithUserAgent(t *testing.T) {
	want := "test-user-agent"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got := r.Header.Get("user-agent")
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("user-agent mismatch: want %v, got %v", got, want)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New("test-token", WithUserAgent(want))

	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.do(req)
	if err != nil {
		t.Fatalf("client.do() failed: %v", err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
}
