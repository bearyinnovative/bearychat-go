package bearychat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	DEFAULT_RTM_API_BASE = "https://rtm.bearychat.com"
)

// RTMClient is used to interactive with BearyChat's RTM api
// and websocket message protocol.
type RTMClient struct {
	// rtm token
	Token string

	// rtm api base, defaults to `https://rtm.bearychat.com`
	APIBase string

	// services
	CurrentTeam *RTMCurrentTeamService
	User        *RTMUserService
	Channel     *RTMChannelService

	httpClient *http.Client
}

type rtmOptSetter func(*RTMClient) error

// enabled services
var services = []rtmOptSetter{
	newRTMCurrentTeamService,
	newRTMUserService,
	newRTMChannelService,
}

// NewRTMClient creates a rtm client.
//
//      client, _ := NewRTMClient(
//              "rtm-token",
//              WithRTMAPIBase("https://rtm.bearychat.com"),
//      )
func NewRTMClient(token string, setters ...rtmOptSetter) (*RTMClient, error) {
	c := &RTMClient{
		Token:   token,
		APIBase: DEFAULT_RTM_API_BASE,

		httpClient: http.DefaultClient,
	}

	for _, setter := range services {
		if err := setter(c); err != nil {
			return nil, err
		}
	}

	for _, setter := range setters {
		if err := setter(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// WithRTMAPIBase can be used to set rtm client's base api.
func WithRTMAPIBase(apiBase string) rtmOptSetter {
	return func(c *RTMClient) error {
		c.APIBase = apiBase
		return nil
	}
}

// WithRTMHTTPClient sets http client.
func WithRTMHTTPClient(httpClient *http.Client) rtmOptSetter {
	return func(c *RTMClient) error {
		c.httpClient = httpClient
		return nil
	}
}

// Do performs an api request.
func (c RTMClient) Do(resource, method string, in, result interface{}) (*http.Response, error) {
	uri, err := addTokenToResourceUri(
		fmt.Sprintf("%s/%s", c.APIBase, resource),
		c.Token,
	)
	if err != nil {
		return nil, err
	}

	// build payload (if any)
	var buf io.ReadWriter
	if in != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(in)
		if err != nil {
			return nil, err
		}
	}

	// build request
	req, err := http.NewRequest(method, uri, buf)
	if err != nil {
		return nil, err
	}
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// parse response
	defer resp.Body.Close()
	response := new(RTMAPIResponse)
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return resp, err
	}

	// request failed
	if resp.StatusCode/100 != 2 || response.Code != 0 {
		return resp, response
	}

	// parse result (if any)
	if result != nil {
		return resp, json.Unmarshal(response.Result, result)
	}

	return resp, nil
}

func (c RTMClient) Get(resource string, result interface{}) (*http.Response, error) {
	return c.Do(resource, "GET", nil, result)
}

func (c RTMClient) Post(resource string, in, result interface{}) (*http.Response, error) {
	return c.Do(resource, "POST", in, result)
}

// Start performs rtm.start
func (c RTMClient) Start() (*User, string, error) {
	userAndWSHost := new(struct {
		User   *User  `json:"user"`
		WSHost string `json:"ws_host"`
	})
	_, err := c.Post("start", nil, userAndWSHost)

	return userAndWSHost.User, userAndWSHost.WSHost, err
}

// RTM api request response
type RTMAPIResponse struct {
	Code        int             `json:"code"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorReason string          `json:"error,omitempty"`
}

func (r *RTMAPIResponse) Error() string {
	return r.ErrorReason
}

func addTokenToResourceUri(resource, token string) (string, error) {
	uri, err := url.Parse(resource)
	if err != nil {
		return "", err
	}

	q := uri.Query()
	q.Set("token", token)
	uri.RawQuery = q.Encode()

	return uri.String(), nil
}
