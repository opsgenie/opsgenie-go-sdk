package contactv2

import (
	"net/url"
	"errors"
)

type Identifier struct {
	ID     string `json:"-"`
	UserID string `json:"-"`
}

func (request *Identifier) GenerateUrl() (string, url.Values, error) {
	if request == nil || len(request.UserID) == 0 {
		return "", nil, errors.New("user ID should be provided")
	}

	baseUrl := "/v2/users/" + request.UserID + "/contacts/"

	if request.ID != "" {
		return baseUrl + request.ID, url.Values{}, nil
	}

	return baseUrl, url.Values{}, nil
}
