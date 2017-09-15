package policyv1

// DisablePolicyRequest is a request to disable policy by id
type DisablePolicyRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey return api key
func (r *DisablePolicyRequest) GetApiKey() string {
	return r.ApiKey
}

// DisablePolicyResponse is a response of disabling policy
type DisablePolicyResponse struct {
	ResponseMeta
	Status string `json:"result"`
}
