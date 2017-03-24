package openapi

import (
	"context"
	"fmt"
	"net/http"
)

type P2P struct {
	ID            *string       `json:"id,omitempty"`
	TeamID        *string       `json:"team_id,omitempty"`
	VChannelID    *string       `json:"vchannel_id,omitempty"`
	Type          *VChannelType `json:"type,omitempty"`
	IsActive      *bool         `json:"is_active"`
	IsMember      *bool         `json:"is_member,omitempty"`
	MemberUserIDs []string      `json:"member_uids,omitempty"`
	LatestTS      *VChannelTS   `json:"latest_ts,omitempty"`
}

type P2PService service

type P2PInfoOptions struct {
	ChannelID string
}

// Info implements `GET /p2p.info`
func (p *P2PService) Info(ctx context.Context, opt *P2PInfoOptions) (*P2P, *http.Response, error) {
	endpoint := fmt.Sprintf("p2p.info?p2p_channel_id=%s", opt.ChannelID)
	req, err := p.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var p2p P2P
	resp, err := p.client.do(ctx, req, &p2p)
	if err != nil {
		return nil, resp, err
	}
	return &p2p, resp, nil
}

// List implements `GET /p2p.list`
func (p *P2PService) List(ctx context.Context) ([]*P2P, *http.Response, error) {
	req, err := p.client.newRequest("GET", "p2p.list", nil)
	if err != nil {
		return nil, nil, err
	}

	var p2p []*P2P
	resp, err := p.client.do(ctx, req, &p2p)
	if err != nil {
		return nil, resp, err
	}
	return p2p, resp, nil
}

type P2PCreateOptions struct {
	UserID string `json:"user_id"`
}

// Create implements `POST /p2p.create`
func (p *P2PService) Create(ctx context.Context, opt *P2PCreateOptions) (*P2P, *http.Response, error) {
	req, err := p.client.newRequest("POST", "p2p.create", opt)
	if err != nil {
		return nil, nil, err
	}

	var p2p P2P
	resp, err := p.client.do(ctx, req, &p2p)
	if err != nil {
		return nil, resp, err
	}
	return &p2p, resp, nil
}
