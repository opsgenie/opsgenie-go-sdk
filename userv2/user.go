package userv2

import "time"

// User contains user data.
type User struct {
	ID           string        `json:"id,omitempty"`
	Blocked      bool          `json:"blocked,omitempty"`
	Verified     bool          `json:"verified,omitempty"`
	UserName     string        `json:"username,omitempty"`
	FullName     string        `json:"fullname,omitempty"`
	Role         Role          `json:"role,omitempty"`
	TimeZone     string        `json:"timeZone,omitempty"`
	Locale       string        `json:"locale,omitempty"`
	UserAddress  UserAddress   `json:"userAddress,omitempty"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
	Details      Details       `json:"details"`
	Tags         Tags          `json:"tags"`
	UserContacts []UserContact `json:"userContacts"`
}
