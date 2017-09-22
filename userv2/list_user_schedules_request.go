package userv2

// ListUserSchedulesRequest is a request for getting list of user schedules.
type ListUserSchedulesRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListUserSchedulesRequest) GetApiKey() string {
	return r.ApiKey
}
