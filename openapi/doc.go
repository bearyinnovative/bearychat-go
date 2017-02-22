// Package openapi implements BearyChat's OpenAPI methods.
//
// For full api doc, please visit https://github.com/bearyinnovative/OpenAPI
//
// Usage:
//
//      import "github.com/bcho/bearychat.go/openapi"
//
// API methods are bound to a API client. To construct a new client:
//
//      client := openapi.NewClient(token)
//
// API methods are grouped by "namespace":
//
//      team, _, err := client.Team.Info()
package openapi
