package teamv2

type CreateTeamRequest struct {
	*Identifier
	ApiKey      string
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Members     []Member `json:"members"`
}

// GetApiKey return api key
func (r *CreateTeamRequest) GetApiKey() string {
	return r.ApiKey
}

type CreateTeamResponse struct {
	ResponseMeta
	Status string `json:"result"`
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}
