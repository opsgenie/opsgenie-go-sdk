package policyv1

import (
	"net/url"
)

const (
	DisableAction = "disable"
	EnableAction = "enable"
)

type Identifier struct {
	ID           string `json:"-"`
	Alias        string `json:"-"`
	TinyID       string `json:"-"`
	StatusAction string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v1/policies"
	if request != nil {
		if len(request.ID) != 0 {
			baseUrl += "/" + request.ID
		}

		if len(request.StatusAction) != 0 {
			baseUrl += "/" + request.StatusAction
		}
	}


	return baseUrl, nil, nil
}
