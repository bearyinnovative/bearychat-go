package openapi

import (
	"context"
	"fmt"
	"net/http"
)

type MessagePin struct {
	ID         *string     `json:"id,omitempty"`
	TeamID     *string     `json:"team_id,omitempty"`
	UID        *string     `json:"uid,omitempty"`
	VchannelID *string     `json:"vchannel_id,omitempty"`
	MessageID  *string     `json:"message_id,omitempty"`
	MessageKey *MessageKey `json:"message_key,omitempty"`
	CreatedAt  *Time       `json:"create_at,omitempty"`
	UpdatedAt  *Time       `json:"update_at,omitempty"`
}

type MessagePinService service

type MessagePinListOptions struct {
	VChannelID string `json:"vchannel_id"`
}

// List implements `GET /message_pin.list`
func (m *MessagePinService) List(ctx context.Context, opt *MessagePinListOptions) ([]*MessagePin, *http.Response, error) {
	endpoint := fmt.Sprintf("message_pin.list?vchannel_id=%s", opt.VChannelID)
	req, err := m.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var messagePins []*MessagePin
	resp, err := m.client.do(ctx, req, &messagePins)
	if err != nil {
		return nil, resp, err
	}

	return messagePins, resp, nil
}

type MessagePinCreateOptions struct {
	VChannelID string     `json:"vchannel_id"`
	MessageKey MessageKey `json:"message_key"`
}

// Create implements `POST /message_pin.create`
func (m *MessagePinService) Create(ctx context.Context, opt *MessagePinCreateOptions) (*MessagePin, *http.Response, error) {
	req, err := m.client.newRequest("POST", "message_pin.create", opt)
	if err != nil {
		return nil, nil, err
	}

	var messagePin MessagePin
	resp, err := m.client.do(ctx, req, &messagePin)
	if err != nil {
		return nil, resp, err
	}
	return &messagePin, resp, nil
}

type MessagePinDeleteOptions struct {
	VChannelID string `json:"vchannel_id"`
	PinID      string `json:"pin_id"`
}

// Delete implements `POST /message_pin.delete`
func (m *MessagePinService) Delete(ctx context.Context, opt *MessagePinDeleteOptions) (*ResponseNoContent, *http.Response, error) {
	req, err := m.client.newRequest("POST", "message_pin.delete", opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := m.client.do(ctx, req, nil)
	if err != nil {
		return nil, resp, err
	}
	return &ResponseNoContent{}, resp, nil
}
