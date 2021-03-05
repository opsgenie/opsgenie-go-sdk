package teamv2

import (
	"errors"
	"net/url"
)

// CreateTeamRequest is a request for creating new team.
type CreateTeamRequest struct {
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Members     []TeamMember `json:"members,omitempty"`
	ApiKey      string       `json:"-"`
}

// GenerateUrl generates url to API endpoint.
func (r *CreateTeamRequest) GenerateUrl() (string, url.Values, error) {
	if r.Name == "" {
		return "", nil, errors.New("Name should be provided for create action")
	}

	return "/v2/teams", nil, nil
}

// GetApiKey returns api key.
func (r *CreateTeamRequest) GetApiKey() string {
	return r.ApiKey
}
