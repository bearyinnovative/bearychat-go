package openapi

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
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

type ChannelCreateOptions struct {
	Name    string  `json:"name"`
	Topic   *string `json:"topic,omitempty"`
	Private *bool   `json:"private,omitempty"`
}

// Create implements `POST /channel.create`
func (c *ChannelService) Create(ctx context.Context, opt *ChannelCreateOptions) (*Channel, *http.Response, error) {
	req, err := c.client.newRequest("POST", "channel.create", opt)
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

type ChannelArchiveOptions struct {
	ChannelID string `json:"channel_id"`
}

// Archive implements `POST /channel.archive`
func (c *ChannelService) Archive(ctx context.Context, opt *ChannelArchiveOptions) (*Channel, *http.Response, error) {
	req, err := c.client.newRequest("POST", "channel.archive", opt)
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

type ChannelUnarchiveOptions struct {
	ChannelID string `json:"channel_id"`
}

// Unarchive implements `POST /channel.unarchive`
func (c *ChannelService) Unarchive(ctx context.Context, opt *ChannelUnarchiveOptions) (*Channel, *http.Response, error) {
	req, err := c.client.newRequest("POST", "channel.unarchive", opt)
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

type ChannelLeaveOptions struct {
	ChannelID string `json:"channel_id"`
}

// Leave implements `POST /channel.leave`
func (c *ChannelService) Leave(ctx context.Context, opt *ChannelLeaveOptions) (*ResponseNoContent, *http.Response, error) {
	req, err := c.client.newRequest("POST", "channel.leave", opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.do(ctx, req, nil)
	if err != nil {
		return nil, resp, err
	}
	return &ResponseNoContent{}, resp, nil
}

type ChannelJoinOptions struct {
	ChannelID string `json:"channel_id"`
}

// Join implements `POST /channel.join`
func (c *ChannelService) Join(ctx context.Context, opt *ChannelJoinOptions) (*Channel, *http.Response, error) {
	req, err := c.client.newRequest("POST", "channel.join", opt)
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

type ChannelInviteOptions struct {
	ChannelID    string `json:"channel_id"`
	InviteUserID string `json:"invite_uid"`
}

// Invite implements `POST /channel.invite`
func (c *ChannelService) Invite(ctx context.Context, opt *ChannelInviteOptions) (*ResponseNoContent, *http.Response, error) {
	req, err := c.client.newRequest("POST", "channel.invite", opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.do(ctx, req, nil)
	if err != nil {
		return nil, resp, err
	}
	return &ResponseNoContent{}, resp, nil
}

type ChannelKickOptions struct {
	ChannelID  string `json:"channel_id"`
	KickUserID string `json:"kick_uid"`
}

// Kick implements `POST /channel.kick`
func (c *ChannelService) Kick(ctx context.Context, opt *ChannelKickOptions) (*ResponseNoContent, *http.Response, error) {
	req, err := c.client.newRequest("POST", "channel.kickout", opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.do(ctx, req, nil)
	if err != nil {
		return nil, resp, err
	}
	return &ResponseNoContent{}, resp, nil
}
