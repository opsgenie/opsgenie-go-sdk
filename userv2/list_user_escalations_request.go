package userv2

// ListUserEscalationsRequest is a request for getting user escalation list.
type ListUserEscalationsRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListUserEscalationsRequest) GetApiKey() string {
	return r.ApiKey
}
