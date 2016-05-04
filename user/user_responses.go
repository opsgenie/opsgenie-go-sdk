package user

// Create escalation response structure
type CreateUserResponse struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Code int `json:"code"`
}

// Update escalation response structure
type UpdateUserResponse struct {
	Status string `json:"status"`
        Code int `json:"code"`
}

// Delete escalation response structure
type DeleteUserResponse struct {
	Status string `json:"status"`
        Code int `json:"code"`
}

// Participant
type Contact struct {
	To string `json:"to,omitempty"`
	Method string `json:"method,omitempty"`
}

// Get escalation structure
type GetUserResponse struct {
	Id string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Fullname string `json:"fullname,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Local string `json:"local,omitempty"`
	State string `json:"state,omitempty"`
	Escalations []string `json:"escalations,omitempty"`
	Schedules []string `json:"schedules,omitempty"`
	Role string `json:"role,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Contacts []Contact `json:"contacts,omitempty"`
}

// List escalations response structure
type ListUsersResponse struct {
	Users []GetUserResponse `json:"users,omitempty"`
}
