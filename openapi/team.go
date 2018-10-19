package openapi

import (
	"context"
	"net/http"
)

type TeamPlan string

type Team struct {
	ID          *string   `json:"id,omitempty"`
	Subdomain   *string   `json:"subdomain,omitempty"`
	Name        *string   `json:"name,omitempty"`
	EmailDomain *string   `json:"email_domain,omitempty"`
	LogoURL     *string   `json:"logo_url,omitempty"`
	Description *string   `json:"description,omitempty"`
	Plan        *TeamPlan `json:"plan,omitempty"`
	Created     *Time     `json:"created,omitempty"`
}

type TeamService service

// Info implements `GET /team.info`
func (t *TeamService) Info(ctx context.Context) (*Team, *http.Response, error) {
	req, err := t.client.newRequest("GET", "team.info", nil)
	if err != nil {
		return nil, nil, err
	}

	var team Team
	resp, err := t.client.do(ctx, req, &team)
	if err != nil {
		return nil, resp, err
	}
	return &team, resp, nil
}
