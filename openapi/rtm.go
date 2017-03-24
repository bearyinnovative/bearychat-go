package openapi

import (
	"context"
	"net/http"
)

type RTMStart struct {
	WebSocketHost *string `json:"ws_host,omitempty"`
	User          *User   `json:"user,omitempty"`
}

type RTMService service

// Start implements `POST /rtm.start`
func (r *RTMService) Start(ctx context.Context) (*RTMStart, *http.Response, error) {
	req, err := r.client.newRequest("POST", "rtm.start", nil)
	if err != nil {
		return nil, nil, err
	}

	var start RTMStart
	resp, err := r.client.do(ctx, req, &start)
	if err != nil {
		return nil, resp, err
	}
	return &start, resp, nil
}
