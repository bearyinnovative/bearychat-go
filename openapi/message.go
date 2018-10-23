package openapi

import (
	"context"
	"fmt"
	"net/http"
)

type MessageKey string

type MessageSubtype string

const (
	MessageSubtypeNormal MessageSubtype = "normal"
	MessageSubtypeInfo                  = "info"
)

type MessageAttachmentImage struct {
	Url *string `json:"url,omitempty"`
}

type MessageAttachment struct {
	Title  *string                  `json:"title,omitempty"`
	Text   *string                  `json:"text,omitempty"`
	Color  *string                  `json:"color,omitempty"`
	Images []MessageAttachmentImage `json:"images,omitempty"`
}

type Reaction struct {
	CreatedTS *VChannelTS `json:"created_ts,omitempty"`
	Reaction  *string     `json:"reaction,omitempty"`
	UIDs      []string    `json:"uids,omitempty"`
}

type Repost struct {
	UID        *string         `json:"uid,omitempty"`
	VchannelID *string         `json:"vchannel_id,omitempty"`
	RobotID    *string         `json:"robot_id,omitempty"`
	CreatedTS  *VChannelTS     `json:"created_ts,omitempty"`
	MessageKey *MessageKey     `json:"message_key,omitempty"`
	ID         *string         `json:"id,omitempty"`
	TeamID     *string         `json:"team_id,omitempty"`
	Subtype    *MessageSubtype `json:"subtype,omitempty"`
	Text       *string         `json:"text,omitempty"`
}

type Message struct {
	Repost          *Repost            `json:"repost,omitempty"`
	Key             *MessageKey        `json:"key,omitempty"`
	Updated         *Time              `json:"updated,omitempty"`
	UID             *string            `json:"uid,omitempty"`
	Created         *Time              `json:"created,omitempty"`
	VchannelID      *string            `json:"vchannel_id,omitempty"`
	ReferKey        *string            `json:"refer_key,omitempty"`
	RobotID         *string            `json:"robot_id,omitempty"`
	Edited          *bool              `json:"edited,omitempty"`
	CreatedTS       *VChannelTS        `json:"created_ts,omitempty"`
	PinID           *string            `json:"pin_id,omitempty"`
	StarID          *string            `json:"star_id,omitempty"`
	ID              *string            `json:"id,omitempty"`
	TeamID          *string            `json:"team_id,omitempty"`
	TextI18n        *map[string]string `json:"text_i18n,omitempty"`
	Reactions       []Reaction         `json:"reactions,omitempty"`
	Subtype         *MessageSubtype    `json:"subtype,omitempty"`
	Text            *string            `json:"text,omitempty"`
	DisableMarkdown *bool              `json:"disable_markdown,omitempty"`
}

type MessageService service

type MessageInfoOptions struct {
	VChannelID string
	Key        MessageKey
}

// Info implements `GET /message.info`
func (m *MessageService) Info(ctx context.Context, opt *MessageInfoOptions) (*Message, *http.Response, error) {
	endpoint := fmt.Sprintf("message.info?vchannel_id=%s&message_key=%s", opt.VChannelID, opt.Key)
	req, err := m.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var message Message
	resp, err := m.client.do(ctx, req, &message)
	if err != nil {
		return nil, resp, err
	}
	return &message, resp, nil
}

type MessageCreateOptions struct {
	VChannelID  string              `json:"vchannel_id"`
	Text        string              `json:"text"`
	Attachments []MessageAttachment `json:"attachments"`
}

// Create implements `POST /message.create`
func (m *MessageService) Create(ctx context.Context, opt *MessageCreateOptions) (*Message, *http.Response, error) {
	if opt.Attachments == nil {
		opt.Attachments = []MessageAttachment{}
	}
	req, err := m.client.newRequest("POST", "message.create", opt)
	if err != nil {
		return nil, nil, err
	}

	var message Message
	resp, err := m.client.do(ctx, req, &message)
	if err != nil {
		return nil, resp, err
	}
	return &message, resp, nil
}

type MessageDeleteOptions struct {
	VChannelID string     `json:"vchannel_id"`
	Key        MessageKey `json:"message_key"`
}

// Delete implements `POST /message.delete`
func (m *MessageService) Delete(ctx context.Context, opt *MessageDeleteOptions) (*ResponseNoContent, *http.Response, error) {
	req, err := m.client.newRequest("POST", "message.delete", opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := m.client.do(ctx, req, nil)
	if err != nil {
		return nil, resp, err
	}
	return &ResponseNoContent{}, resp, nil
}

type MessageUpdateTextOptions struct {
	VChannelID string     `json:"vchannel_id"`
	Key        MessageKey `json:"message_key"`
	Text       string     `json:"text"`
}

//UpdateText implements `POST /message.update_text`
func (m *MessageService) UpdateText(ctx context.Context, opt *MessageUpdateTextOptions) (*Message, *http.Response, error) {
	req, err := m.client.newRequest("PATCH", "message.update_text", opt)
	if err != nil {
		return nil, nil, err
	}

	var message Message
	resp, err := m.client.do(ctx, req, &message)
	if err != nil {
		return nil, resp, err
	}
	return &message, resp, nil
}

type MessageForwardOptions struct {
	VChannelID   string     `json:"vchannel_id"`
	Key          MessageKey `json:"message_key"`
	ToVChannelID string     `json:"to_vchannel_id"`
}

// Forward implements `POST /message.forward`
func (m *MessageService) Forward(ctx context.Context, opt *MessageForwardOptions) (*Message, *http.Response, error) {
	req, err := m.client.newRequest("POST", "message.forward", opt)
	if err != nil {
		return nil, nil, err
	}

	var message Message
	resp, err := m.client.do(ctx, req, &message)
	if err != nil {
		return nil, resp, err
	}
	return &message, resp, nil
}
