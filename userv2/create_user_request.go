package userv2

import "net/url"

// CreateUserRequest is a request for creating new user.
type CreateUserRequest struct {
	UserName           string      `json:"username,omitempty"`
	FullName           string      `json:"fullName,omitempty"`
	Role               Role        `json:"role,omitempty"`
	SkypeUsername      string      `json:"skypeUsername,omitempty"`
	UserAddress        UserAddress `json:"userAddress,omitempty"`
	Tags               Tags        `json:"tags,omitempty"`
	Details            Details     `json:"details,omitempty"`
	Timezone           string      `json:"timezone,omitempty"`
	Locale             string      `json:"locale,omitempty"`
	InvitationDisabled bool        `json:"invitationDisabled,omitempty"`
	ApiKey             string
}

// GenerateUrl generates url to API endpoint.
func (r *CreateUserRequest) GenerateUrl() (string, url.Values, error) {
	return "/v2/users", nil, nil
}

// GetApiKey returns api key.
func (r *CreateUserRequest) GetApiKey() string {
	return r.ApiKey
}
