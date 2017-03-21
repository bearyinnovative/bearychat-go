package openapi

import (
	"context"
	"fmt"
	"net/http"
)

type Channel struct {
	ID            *string       `json:"id,omitempty"`
	TeamID        *string       `json:"team_id,omitempty"`
	VChannelID    *string       `json:"vchannel_id,omitempty"`
	UserID        *string       `json:"uid,omitempty"`
	Name          *string       `json:"name,omitempty"`
	Type          *VChannelType `json:"type,omitempty"`
	Private       *bool         `json:"private,omitempty"`
	General       *bool         `json:"general,omitempty"`
	Topic         *string       `json:"topic,omitempty"`
	IsMember      *bool         `json:"is_member,omitempty"`
	IsActive      *bool         `json:"is_active,omitempty"`
	MemberUserIDs []string      `json:"member_uids,omitempty"`
	LatestTS      *VChannelTS   `json:"latest_ts,omitempty"`
}

type ChannelService service

type ChannelInfoOptions struct {
	ChannelID string
}

// Info implements `GET /channel.info`
func (c *ChannelService) Info(ctx context.Context, opt *ChannelInfoOptions) (*Channel, *http.Response, error) {
	endpoint := fmt.Sprintf("channel.info?channel_id=%s", opt.ChannelID)
	req, err := c.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var channel Channel
	resp, err := c.client.do(ctx, req, &channel)
	if err != nil {
		return nil, resp, err
	}
	return &channel, resp, nil
}

// List implements `GET /channel.list`
func (c *ChannelService) List(ctx context.Context) ([]*Channel, *http.Response, error) {
	req, err := c.client.newRequest("GET", "channel.list", nil)
	if err != nil {
		return nil, nil, err
	}

	var channels []*Channel
	resp, err := c.client.do(ctx, req, &channels)
	if err != nil {
		return nil, resp, err
	}
	return channels, resp, nil
}
