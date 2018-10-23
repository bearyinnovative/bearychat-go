package openapi

import (
	"context"
	"net/http"
)

type Sticker struct {
	URL    *string `json:"url,omitempty"`
	Name   *string `json:"name,omitempty"`
	Width  *int    `json:"width,omitempty"`
	Height *int    `json:"height,omitempty"`
}

type StickerPack struct {
	PackName *string    `json:"pack,omitempty"`
	Stickers []*Sticker `json:"stickers,omitempty"`
}

type StickerService service

// List implements `GET /sticker.list`
func (s *StickerService) List(ctx context.Context) ([]*StickerPack, *http.Response, error) {
	req, err := s.client.newRequest("GET", "sticker.list", nil)
	if err != nil {
		return nil, nil, err
	}

	var stickers []*StickerPack
	resp, err := s.client.do(ctx, req, &stickers)
	if err != nil {
		return nil, resp, err
	}
	return stickers, resp, nil
}
