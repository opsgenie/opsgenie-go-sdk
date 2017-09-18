package teamv2

type GetTeamRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey return api key
func (r *GetTeamRequest) GetApiKey() string {
	return r.ApiKey
}

type GetTeamResponse struct {
	ResponseMeta
	Team struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Members     []Member
	} `json:"data"`
}
