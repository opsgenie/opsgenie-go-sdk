package teamv2

import (
	"net/url"
)

// DeleteTeamRequest is a request for deleting user.
type DeleteTeamRequest struct {
	*Identifier
	ApiKey string
}

// GenerateUrl generates url to API endpoint.
func (r *DeleteTeamRequest) GenerateUrl() (string, url.Values, error) {
	baseUrl, _, err := r.Identifier.GenerateUrl()

	if err != nil {
		return "", nil, err
	}

	return baseUrl, nil, err
}

// GetApiKey returns api key.
func (r *DeleteTeamRequest) GetApiKey() string {
	return r.ApiKey
}
