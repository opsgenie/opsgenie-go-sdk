package escalationv2

type ListEscalationRequest struct {
	*Identifier
	ApiKey
}

type ListEscalationResponse struct {
	ResponseMeta
	Escalations []Escalation `json:"data"`
}
