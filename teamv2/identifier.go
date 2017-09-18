package teamv2

import (
	"net/url"
)

type Identifier struct {
	ID     string `json:"-"`
	Alias  string `json:"-"`
	TinyID string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/teams"

	return baseUrl, nil, nil
}
