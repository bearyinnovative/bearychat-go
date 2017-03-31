package openapi

import (
	"context"
	"net/http"
)

func uintp(x uint) *uint { return &x }

var (
	// Shorthand function for passing limit value as pointer.
	MessageQueryWithLimit = uintp
	// Shorthand function for passing forward value as pointer.
	MessageQueryWithForward = uintp
	// Shorthand function for passing backward value as pointer.
	MessageQueryWithBackward = uintp
)

// TODO(hbc): introduce a query builder or map literal
type MessageQuery struct {
	Latest *MessageQueryByLatest `json:"latest,omitempty"`
	Since  *MessageQueryBySince  `json:"since,omitempty"`
	Window *MessageQueryByWindow `json:"window,omitempty"`
}

type MessageQueryByLatest struct {
	Limit *uint `json:"limit,omitempty"`
}

type MessageQueryBySince struct {
	SinceKey *MessageKey `json:"key,omitempty"`
	SinceTS  *VChannelTS `json:"ts,omitempty"`
	Forward  *uint       `json:"forward,omitempty"`
	Backward *uint       `json:"backward,omitempty"`
}

type MessageQueryByWindow struct {
	FromKey  *MessageKey `json:"from_key,omitempty"`
	ToKey    *MessageKey `json:"to_key,omitempty"`
	FromTS   *VChannelTS `json:"from_ts,omitempty"`
	ToTS     *VChannelTS `json:"to_ts,omitempty"`
	Forward  *uint       `json:"forward,omitempty"`
	Backward *uint       `json:"backward,omitempty"`
}

type MessageQueryOptions struct {
	VChannelID string        `json:"vchannel_id"`
	Query      *MessageQuery `json:"query"`
}

type MessageQueryResult struct {
	Messages []*Message `json:"messages"`
}

// Query implements `POST /message.query`
func (m *MessageService) Query(ctx context.Context, opt *MessageQueryOptions) (*MessageQueryResult, *http.Response, error) {
	req, err := m.client.newRequest("POST", "message.query", opt)
	if err != nil {
		return nil, nil, err
	}

	var rv MessageQueryResult
	resp, err := m.client.do(ctx, req, &rv)
	if err != nil {
		return nil, resp, err
	}
	return &rv, resp, nil
}
