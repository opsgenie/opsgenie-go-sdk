package schedulev2

type GetScheduleRequest struct {
	*Identifier
	ApiKey
}

type GetScheduleResponse struct {
	ResponseMeta
	Schedule Schedule `json:"data"`
}
