package openapi

import (
	"net/http"

	"golang.org/x/net/context"
)

type Emoji struct {
	ID      *string `json:"id,omitempty"`
	UserID  *string `json:"uid,omitempty"`
	TeamID  *string `json:"team_id,omitempty"`
	Name    *string `json:"name,omitempty"`
	URL     *string `json:"url,omitempty"`
	Created *Time   `json:"created,omitempty"`
	Updated *Time   `json:"updated,omitempty"`
}

type EmojiService service

// List implements `GET /emoji.list`
func (e *EmojiService) List(ctx context.Context) ([]*Emoji, *http.Response, error) {
	req, err := e.client.newRequest("GET", "emoji.list", nil)
	if err != nil {
		return nil, nil, err
	}

	var emojis []*Emoji
	resp, err := e.client.do(ctx, req, &emojis)
	if err != nil {
		return nil, resp, err
	}
	return emojis, resp, nil
}
