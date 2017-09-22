package userv2

// DeleteUserRequest is a request for deleting user.
type DeleteUserRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *DeleteUserRequest) GetApiKey() string {
	return r.ApiKey
}
