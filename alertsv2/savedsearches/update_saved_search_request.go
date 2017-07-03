package alertsv2

import (
	"net/url"
	"errors"
)

type UpdateSavedSearchRequest struct {
	ID          string `json:"-"`
	Name        string `json:"-"`
	NewName     string `json:"name,omitempty"`
	Query       string `json:"query,omitempty"`
	Owner       User `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
	Teams       []Team `json:"teams,omitempty"`
	ApiKey      string `json:"-"`
}

func (r *UpdateSavedSearchRequest) GetApiKey() string {
	return r.ApiKey
}

func (r *UpdateSavedSearchRequest) GenerateUrl() (string, url.Values, error) {
	if r.ID != "" {
		return "/v2/alerts/saved-searches/" + r.ID, nil, nil
	}

	if r.Name != "" {
		params := url.Values{}
		params.Add("identifierType", "name")
		return "/v2/alerts/saved-searches/" + r.Name, params, nil
	}

	return "", nil, errors.New("ID or Name should be provided")
}
