package userv2

// ListUserTeamsRequest is a request for getting list of user teams.
type ListUserTeamsRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListUserTeamsRequest) GetApiKey() string {
	return r.ApiKey
}
