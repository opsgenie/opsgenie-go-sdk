package teamv2

import "net/url"

// GetTeamRequest is a request for getting team.
type GetTeamRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *GetTeamRequest) GetApiKey() string {
	return r.ApiKey
}

// GenerateUrl generates API url using specified attributes of identifier.
func (r *GetTeamRequest) GenerateUrl() (string, url.Values, error) {
	baseUrl, params, err := r.Identifier.GenerateUrl()

	if err != nil {
		return "", nil, err
	}

	return baseUrl, params, nil
}
