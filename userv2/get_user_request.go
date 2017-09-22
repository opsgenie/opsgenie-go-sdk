package userv2

// GetUserRequest is a request for getting user.
type GetUserRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *GetUserRequest) GetApiKey() string {
	return r.ApiKey
}
