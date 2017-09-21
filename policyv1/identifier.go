package policyv1

import (
	"net/url"
)

const (
	DisableAction = "disable"
	EnableAction  = "enable"
)

type Identifier struct {
	ID           string `json:"-"`
	Name         string `json:"name"`
	StatusAction string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v1/policies"
	if request != nil {
		if len(request.Name) != 0 {
			baseUrl += "/" + request.Name
		}

		if len(request.ID) != 0 {
			baseUrl += "/" + request.ID
		}

		if len(request.StatusAction) != 0 {
			baseUrl += "/" + request.StatusAction
		}
	}

	return baseUrl, nil, nil
}
