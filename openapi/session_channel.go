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

type SessionChannelCreateOptions struct {
	Name          *string  `json:"name,omitempty"`
	MemberUserIDs []string `json:"member_uids"`
}

// Create implements `POST /session_channel.create`
func (s *SessionChannelService) Create(ctx context.Context, opt *SessionChannelCreateOptions) (*SessionChannel, *http.Response, error) {
	req, err := s.client.newRequest("POST", "session_channel.create", opt)
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

type SessionChannelArchiveOptions struct {
	ChannelID string `json:"session_channel_id"`
}

// Archive implements `POST /session_channel.archive`
func (s *SessionChannelService) Archive(ctx context.Context, opt *SessionChannelArchiveOptions) (*SessionChannel, *http.Response, error) {
	req, err := s.client.newRequest("POST", "session_channel.archive", opt)
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

type SessionChannelConvertOptions struct {
	ChannelID string `json:"session_channel_id"`
	Name      string `json:"name"`
	Private   *bool  `json:"private,omitempty"`
}

// ConvertToChannel implements `POST /session_channel.convert_to_channel`
func (s *SessionChannelService) ConvertToChannel(ctx context.Context, opt *SessionChannelConvertOptions) (*Channel, *http.Response, error) {
	req, err := s.client.newRequest("POST", "session_channel.convert_to_channel", opt)
	if err != nil {
		return nil, nil, err
	}

	var channel Channel
	resp, err := s.client.do(ctx, req, &channel)
	if err != nil {
		return nil, resp, err
	}
	return &channel, resp, nil
}

type SessionChannelLeaveOptions struct {
	ChannelID string `json:"session_channel_id"`
}

// Leave implements `POST /session_channel.leave`
func (s *SessionChannelService) Leave(ctx context.Context, opt *SessionChannelLeaveOptions) (*ResponseNoContent, *http.Response, error) {
	req, err := s.client.newRequest("POST", "session_channel.leave", opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.do(ctx, req, nil)
	if err != nil {
		return nil, resp, err
	}
	return &ResponseNoContent{}, resp, nil
}
