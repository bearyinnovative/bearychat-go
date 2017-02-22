package openapi

import (
	"net/http"
	"net/url"
)

var defaultBaseURL = "https://api.bearychat.com/v1/"

// Client interacts with BearyChat's API.
type Client struct {
	// HTTP client for calling the API.
	// Use http.DefaultClient by default.
	httpClient *http.Client

	// Base URL for API requests. Defaults to BearyChat's OpenAPI host.
	// BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	// Access token for the client.
	Token string

	// Shared services holder to reduce real service allocating.
	base service
}

type service struct {
	client *Client
}

type clientOpt func(c *Client)

// NewClientWithBaseURL binds BaseURL to client.
func NewClientWithBaseURL(u *url.URL) clientOpt {
	return func(c *Client) {
		c.BaseURL = u
	}
}

// NewClientWithHTTPClient binds http client to client.
func NewClientWithHTTPClient(httpClient *http.Client) clientOpt {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// NewClient constructs a client with given access token.
// Other settings can set via clientOpt functions.
func NewClient(token string, opts ...clientOpt) *Client {
	c := &Client{
		Token: token,
	}

	for _, o := range opts {
		o(c)
	}

	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	if c.BaseURL == nil {
		baseURL, _ := url.Parse(defaultBaseURL)
		c.BaseURL = baseURL
	}

	c.base.client = c

	return c
}
