package teamv2

// GetTeamResponse is a response of getting team result.
type GetTeamResponse struct {
	Team Team `json:"data"`
	ResponseMeta
}
