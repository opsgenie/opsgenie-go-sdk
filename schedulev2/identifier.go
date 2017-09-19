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

	if request != nil {
		if len(request.ID) != 0 {
			baseUrl += request.ID
		} else if len(request.Name) != 0 {
			baseUrl += request.Name
		}
	}

	return baseUrl, url.Values{}, nil
}
