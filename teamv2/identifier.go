package teamv2

import (
	"errors"
	"net/url"
)

// Identifier defined the set of attributes for identification team.
type Identifier struct {
	ID   string `json:"-"`
	Name string `json:"-"`
}

// GenerateUrl generates API url using specified attributes of identifier.
func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	baseUrl := "/v2/teams"

	identifierType := "id"
	if len(request.ID) > 0 {
		baseUrl += "/" + request.ID
	} else if len(request.Name) > 0 {
		baseUrl += "/" + request.Name
		identifierType = "name"
	} else {
		return "", nil, errors.New("name or id of the team should be provided")
	}

	return baseUrl, url.Values{"identifierType": []string{identifierType}}, nil
}
