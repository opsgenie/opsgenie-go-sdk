package escalationv2

import (
	"net/url"
)

type Identifier struct {
	ID string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/escalations/"

	if request != nil {
		baseUrl += request.ID
	}

	return baseUrl, url.Values{}, nil
}
