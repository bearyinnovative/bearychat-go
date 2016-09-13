package bearychat

import (
	"errors"
	"testing"
)

const (
	testRTMToken = "foobar"
)

func TestNewRTMClient(t *testing.T) {
	c, err := NewRTMClient(testRTMToken)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	if c.Token != testRTMToken {
		t.Errorf("unexpected token: %s", c.Token)
	}

	if c.APIBase != DEFAULT_RTM_API_BASE {
		t.Errorf("should use default rtm api base: %s", c.APIBase)
	}
}

func TestNewRTMClient_error(t *testing.T) {
	newError := errors.New("error")
	_, err := NewRTMClient(testRTMToken, func(c *RTMClient) error {
		return newError
	})
	if err != newError {
		t.Errorf("should return error: %+v", err)
	}
}

func TestNewRTMClient_WithRTMAPIBase(t *testing.T) {
	apiBase := "http://foo.bar"
	c, err := NewRTMClient(testRTMToken, WithRTMAPIBase(apiBase))
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	if c.APIBase != apiBase {
		t.Errorf("should set api base: %s", c.APIBase)
	}
}

func TestNewRTMClient_WithRTMHTTPClient(t *testing.T) {
	c, err := NewRTMClient(testRTMToken, WithRTMHTTPClient(nil))
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	if c.httpClient != nil {
		t.Errorf("should set http client: %+v", c.httpClient)
	}
}

func testAddTokenToResourceUri(t *testing.T) {
	u, err := addTokenToResourceUri("http://foobar.com", "foobar")
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}
	if u != "http://foobar.com?token=foobar" {
		t.Errorf("unexpected resource uri: %s", u)
	}
}
