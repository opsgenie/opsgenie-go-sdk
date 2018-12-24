package teamv2

import "net/url"

// UpdateTeamRequest is a request for updating team.
type UpdateTeamRequest struct {
	*Identifier
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Members     []TeamMember `json:"members,omitempty"`
	ApiKey      string       `json:"-"`
}

// GetApiKey returns api key.
func (r *UpdateTeamRequest) GetApiKey() string {
	return r.ApiKey
}

// GenerateUrl generates url to API endpoint.
func (r *UpdateTeamRequest) GenerateUrl() (string, url.Values, error) {
	baseUrl, _, err := r.Identifier.GenerateUrl()

	if err != nil {
		return "", nil, err
	}

	return baseUrl, nil, err
}
