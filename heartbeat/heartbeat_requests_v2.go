package heartbeat

import (
	"errors"
	"net/url"
)

type PingHeartbeatRequest struct {
	Name   string
	APIKey string
}

func (r *PingHeartbeatRequest) GetApiKey() string {
	return r.APIKey
}

func (r *PingHeartbeatRequest) GenerateUrl() (string, url.Values, error) {
	if r.Name == "" {
		return "", nil, errors.New("Name should be provided")
	}
	return "/v2/heartbeats/" + r.Name + "/ping", nil, nil
}

type ListHeartbeatsRequestV2 struct {
	APIKey string
}

func (r *ListHeartbeatsRequestV2) GetApiKey() string {
	return r.APIKey
}

func (r *ListHeartbeatsRequestV2) GenerateUrl() (string, url.Values, error) {
	return "/v2/heartbeats", nil, nil
}
