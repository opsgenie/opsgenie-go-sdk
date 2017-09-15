package policyv1

// GetPolicyRequest is a request struct to load policy by id
type GetPolicyRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key
func (r *GetPolicyRequest) GetApiKey() string {
	return r.ApiKey
}

// GetPolicyResponse is a response of getting policy
type GetPolicyResponse struct {
	ResponseMeta
	Policy GetPolicyResult `json:"data"`
}

// GetPolicyResult is a result struct of getting policy
type GetPolicyResult struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type PolicyType `json:"type"`
	PolicyDescription string `json:"policyDescription"`
	Enabled bool `json:"enabled"`
	Filter Filter `json:"filter"`
	DeduplicationActionType string `json:"deduplication_action_type"`
	Count int `json:"count"`
}
