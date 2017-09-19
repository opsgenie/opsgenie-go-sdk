package schedulev2

import (
	"net/url"
)

type Identifier struct {
	ID   string `json:"-"`
	Name string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/schedules/"

	if request != nil && request.ID != "" {
		return baseUrl + request.ID, url.Values{}, nil
	}

	params := url.Values{}

	return baseUrl, params, nil
}
