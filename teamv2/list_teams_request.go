package teamv2

import (
	"net/url"
)

type ListTeamsRequest struct {
	ApiKey string
}

// GenerateUrl generates url to API endpoint.
func (r *ListTeamsRequest) GenerateUrl() (string, url.Values, error) {
	return "/v2/teams", nil, nil
}

// GetApiKey returns api key.
func (r *ListTeamsRequest) GetApiKey() string {
	return r.ApiKey
}
