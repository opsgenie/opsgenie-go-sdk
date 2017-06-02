package alertsv2

import (
	"net/url"
	"errors"
)

type GetRequestStatusRequest struct {
	RequestID string `json:"requestId,omitempty"`
	ApiKey    string
}

func (r *GetRequestStatusRequest) GetApiKey() string {
	return r.ApiKey
}

func (r *GetRequestStatusRequest) GenerateUrl() (string, url.Values, error) {
	if r.RequestID != "" {
		return "/v2/alerts/requests/" + r.RequestID, nil, nil
	}

	return "", nil, errors.New("RequestID should be provided")
}