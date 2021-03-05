package alertsv2

import (
	"net/url"
)

type CountAlertRequest struct {
	Query                string
	SearchIdentifier     string
	SearchIdentifierType SearchIdentifierType
	ApiKey               string
}

func (r *CountAlertRequest) GetApiKey() string {
	return r.ApiKey
}

func (request *CountAlertRequest) GenerateUrl() (string, url.Values, error) {
	params := url.Values{}

	if request.Query != "" {
		params.Add("query", request.Query)
	}

	if request.SearchIdentifier != "" {
		params.Add("searchIdentifier", request.SearchIdentifier)
	}

	if request.SearchIdentifierType != "" {
		params.Add("searchIdentifierType", string(request.SearchIdentifierType))
	}

	return "/v2/alerts/count", params, nil
}
