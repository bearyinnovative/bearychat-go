package openapi

import (
	"context"
	"net/http"
)

type Meta struct {
	Version *string `json:"version,omitempty"`
}

type MetaService service

// Get implements `GET /meta`
func (m *MetaService) Get(ctx context.Context) (*Meta, *http.Response, error) {
	req, err := m.client.newRequest("GET", "meta", nil)
	if err != nil {
		return nil, nil, err
	}

	var meta Meta
	resp, err := m.client.do(ctx, req, &meta)
	if err != nil {
		return nil, resp, err
	}
	return &meta, resp, nil
}
