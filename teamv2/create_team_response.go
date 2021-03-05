package teamv2

// CreateTeamResponse is a response of creating user result.
type CreateTeamResponse struct {
	ResponseMeta
	ActionResult
	TeamMeta `json:"data"`
}
