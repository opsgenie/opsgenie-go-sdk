package teamv2

// ListUsersResponse is a response with list of teams.
type ListTeamsResponse struct {
	Teams []Team `json:"data"`
	ResponseMeta
}
