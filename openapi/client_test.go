package openapi

import (
	"net/http"
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	token := "foobar"
	client := NewClient(token)

	if client.httpClient != http.DefaultClient {
		t.Errorf("should use http.DefaultClient by default")
	}
	if client.BaseURL.String() != defaultBaseURL {
		t.Errorf("should use defaultBaseURL by default")
	}
	if client.Token != token {
		t.Errorf("unexpected token: %s", client.Token)
	}
}

func TestNewClient_NewClientWithBaseURL(t *testing.T) {
	u, _ := url.Parse("http://foobar/")
	client := NewClient("foobar", NewClientWithBaseURL(u))

	if client.BaseURL != u {
		t.Errorf("unexpected BaseURL: %s", client.BaseURL)
	}
}

func TestNewClient_NewClientWithHTTPClient(t *testing.T) {
	httpClient := &http.Client{}
	client := NewClient("foobar", NewClientWithHTTPClient(httpClient))

	if client.httpClient != httpClient {
		t.Errorf("unexpected httpClient: %+v", client.httpClient)
	}
}
