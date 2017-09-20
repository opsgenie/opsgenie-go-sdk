package schedulev2

type ListScheduleRequest struct {
	ApiKey
	*Identifier
}

type ListScheduleResponse struct {
	ResponseMeta
	Schedules []Schedule `json:"data"`
	Expandable []string `json:"expandable"`
}
