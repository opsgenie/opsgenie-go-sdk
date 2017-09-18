package teamv2

import (
	"net/url"
	"errors"
)

type Identifier struct {
	ID   string `json:"-"`
	Name string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/teams"

	if request.ID != "" {
		return baseUrl + request.ID, url.Values{}, nil
	}

	params := url.Values{}

	if request.Name != "" {
		params.Set("identifierType", "name")
		return baseUrl + request.Name, params, nil
	}

	return "", nil, errors.New("ID or Name should be provided")
}
