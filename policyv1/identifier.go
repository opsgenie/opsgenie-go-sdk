package policyv1

import (
	"net/url"
)

type Identifier struct {
	ID     string `json:"-"`
	Alias  string `json:"-"`
	TinyID string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v1/policies"
	if request != nil && len(request.ID) != 0 {
		baseUrl += "/" + request.ID
	}

	return baseUrl, nil, nil
}