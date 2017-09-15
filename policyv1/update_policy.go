package policyv1

// UpdatePolicyRequest is a request to update policy
type UpdatePolicyRequest struct {
	*Identifier
	ApiKey            string
	Name              string            `json:"name"`
	Type              PolicyType        `json:"type"`
	PolicyDescription string            `json:"policyDescription"`
	Enabled           bool              `json:"enabled"`
	Filter            Filter            `json:"filter"`
	TimeRestrictions  *TimeRestrictions `json:"timeRestrictions,omitempty"`
	Duration          Duration          `json:"duration"`
}

// GetApiKey return api key
func (r *UpdatePolicyRequest) GetApiKey() string {
	return r.ApiKey
}

// UpdatePolicyResponse is a response of updating policy
type UpdatePolicyResponse struct {
	ResponseMeta
	Status string `json:"result"`
}