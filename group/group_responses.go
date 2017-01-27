package group

// Create group response structure
type CreateGroupResponse struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Code int `json:"code"`
}

// Update group response structure
type UpdateGroupResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`
}

// Delete group response structure
type DeleteGroupResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`
}

// Get group response structure
type GetGroupResponse struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Users []string `json:"users,omitempty"`
}

// List groups response structure
type ListGroupsResponse struct {
	Groups []GetGroupResponse `json:"groups,omitempty"`
}
