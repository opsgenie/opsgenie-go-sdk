package userv2

// UpdateUserRequest is a request for updating user.
type UpdateUserRequest struct {
	*Identifier
	FullName      string       `json:"fullName,omitempty"`
	Role          *Role        `json:"role,omitempty"`
	SkypeUsername string       `json:"skypeUsername,omitempty"`
	UserAddress   *UserAddress `json:"userAddress,omitempty"`
	Tags          *Tags        `json:"tags,omitempty"`
	Details       *Details     `json:"details,omitempty"`
	Timezone      string       `json:"timezone,omitempty"`
	Locale        string       `json:"locale,omitempty"`
	ApiKey        string
}

// GetApiKey returns api key.
func (r *UpdateUserRequest) GetApiKey() string {
	return r.ApiKey
}
