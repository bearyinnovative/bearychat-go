package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/context"
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

	Meta           *MetaService
	Team           *TeamService
	User           *UserService
	Channel        *ChannelService
	SessionChannel *SessionChannelService
	Message        *MessageService
	P2P            *P2PService
	Emoji          *EmojiService
	Sticker        *StickerService
	RTM            *RTMService
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
	c.Meta = (*MetaService)(&c.base)
	c.Team = (*TeamService)(&c.base)
	c.User = (*UserService)(&c.base)
	c.Channel = (*ChannelService)(&c.base)
	c.SessionChannel = (*SessionChannelService)(&c.base)
	c.Message = (*MessageService)(&c.base)
	c.P2P = (*P2PService)(&c.base)
	c.Emoji = (*EmojiService)(&c.base)
	c.Sticker = (*StickerService)(&c.base)
	c.RTM = (*RTMService)(&c.base)

	return c
}

// newRequest creates an API request. API method should specified without a leading slash.
// If specified, the value pointed to body is JSON encoded and included as the request body.
func (c *Client) newRequest(requestMethod, apiMethod string, body interface{}) (*http.Request, error) {
	m, err := url.Parse(apiMethod)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(m)
	q := u.Query()
	q.Set("token", c.Token)
	u.RawQuery = q.Encode()

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(requestMethod, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// do sends an API request and returns the API response. The API response is JSON decoded and
// stored in the value pointed to v. If v implements the io.Writer interface, the raw response body
// will be written to v, without attempting to first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out, ctx.Err() will be returned.
func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		// try to use context's error
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

// ErrorResponse represents errors caused by an API request.
type ErrorResponse struct {
	// HTTP response that caused this error
	Response    *http.Response
	ErrorCode   int    `json:"code"`
	ErrorReason string `json:"error"`
}

func (r ErrorResponse) Error() string {
	return fmt.Sprintf(
		"%v %v: %d %d %s",
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Response.StatusCode,
		r.ErrorCode,
		r.ErrorReason,
	)
}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errResponse)
	}

	// TODO(hbc): handle ratelimit error

	return errResponse
}

const timeLayout = "2006-01-02T15:04:05-0700"

// Time with custom JSON format.
type Time struct {
	time.Time
}

func (t Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	t.Time, err = time.Parse(timeLayout, s)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(timeLayout))), nil
}

type ResponseOK struct {
	Code *int `json:"code,omitempty"`
}

type ResponseNoContent struct{}
