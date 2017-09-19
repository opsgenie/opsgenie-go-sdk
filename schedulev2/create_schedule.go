package schedulev2

type CreateScheduleRequest struct {
	*Identifier
	ApiKey

	Name        string     `json:"name"`
	Description string     `json:"description"`
	Timezone    string     `json:"timezone"`
	Enabled     bool       `json:"enabled"`
	OwnerTeam   OwnerTeam  `json:"ownerTeam"`
	Rotations   []Rotation `json:"rotations"`
}

type CreateScheduleResponse struct {
	ResponseMeta
	Result CreateScheduleResult `json:"data"`
}

type CreateScheduleResult struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
}
