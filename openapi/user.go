package openapi

import (
	"context"
	"fmt"
	"net/http"
)

type UserType string

type UserRole string

type UserAvatar struct {
	Small  *string `json:"small,omitempty"`
	Medium *string `json:"medium,omitempty"`
	Large  *string `json:"large,omitempty"`
}

type UserProfile struct {
	Bio      *string `json:"bio,omitempty"`
	Position *string `json:"position,omitempty"`
	Skype    *string `json:"sykpe,omitempty"`
}

type User struct {
	ID       *string      `json:"id,omitempty"`
	TeamID   *string      `json:"team_id,omitempty"`
	Email    *string      `json:"email,omitempty"`
	Name     *string      `json:"name,omitempty"`
	FullName *string      `json:"full_name,omitempty"`
	Type     *UserType    `json:"type,omitempty"`
	Role     *UserRole    `json:"role,omitempty"`
	Avatars  *UserAvatar  `json:"avatars,omitempty"`
	Profile  *UserProfile `json:"profile,omitempty"`
	Inactive *bool        `json:"inactive,omitempty"`
	Created  *Time        `json:"created,omitempty"`
}

type UserService service

type UserInfoOptions struct {
	UserID string
}

// Info implements `GET /user.info`
func (u *UserService) Info(ctx context.Context, opt *UserInfoOptions) (*User, *http.Response, error) {
	endpoint := fmt.Sprintf("user.info?user_id=%s", opt.UserID)
	req, err := u.client.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := u.client.do(ctx, req, &user)
	if err != nil {
		return nil, resp, err
	}
	return &user, resp, nil
}

// List implements `GET /user.list`
func (u *UserService) List(ctx context.Context) ([]*User, *http.Response, error) {
	req, err := u.client.newRequest("GET", "user.list", nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := u.client.do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}
	return users, resp, nil
}

// Me implements `GET /user.me`
func (u *UserService) Me(ctx context.Context) (*User, *http.Response, error) {
	req, err := u.client.newRequest("GET", "user.me", nil)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := u.client.do(ctx, req, &user)
	if err != nil {
		return nil, resp, err
	}
	return &user, resp, nil
}
