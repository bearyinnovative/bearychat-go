package openapi

import (
	"context"
	"fmt"
	"net/http"
)

type SessionChannel struct {
	ID            *string       `json:"id,omitempty"`
	TeamID        *string       `json:"team_id,omitempty"`
	VChannelID    *string       `json:"vchannel_id,omitempty"`
	Name          *string       `json:"name,omitempty"`
	Type          *VChannelType `json:"type,omitempty"`
	IsMember      *bool         `json:"is_member,omitempty"`
	IsActive      *bool         `json:"is_active,omitempty"`
	MemberUserIDs []string      `json:"member_uids,omitempty"`
	LatestTS      *VChannelTS   `json:"latest_ts,omitempty"`
}

type SessionChannelService service

type SessionChannelInfoOptions struct {
	ChannelID string
}

// Info implements `GET /session_channel.info`
func (s *SessionChannelService) Info(ctx context.Context, opt *SessionChannelInfoOptions) (*SessionChannel, *http.Response, error) {
	endpoint := fmt.Sprintf("session_channel.info?session_channel_id=%s", opt.ChannelID)
	req, err := s.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var channel SessionChannel
	resp, err := s.client.do(ctx, req, &channel)
	if err != nil {
		return nil, resp, err
	}
	return &channel, resp, nil
}

// List implements `GET /session_channel.list`
func (s *SessionChannelService) List(ctx context.Context) ([]*SessionChannel, *http.Response, error) {
	req, err := s.client.newRequest("GET", "session_channel.list", nil)
	if err != nil {
		return nil, nil, err
	}

	var channels []*SessionChannel
	resp, err := s.client.do(ctx, req, &channels)
	if err != nil {
		return nil, resp, err
	}
	return channels, resp, nil
}
