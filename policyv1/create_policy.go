package policyv1

// CreatePolicyRequest is a request to create new policy
type CreatePolicyRequest struct {
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
func (r *CreatePolicyRequest) GetApiKey() string {
	return r.ApiKey
}

// CreatePolicyResponse is a response of creating new policy
type CreatePolicyResponse struct {
	ResponseMeta
	Policy CreatePolicyResult `json:"data"`
}

// CreatePolicyResult is struct, which contains only base attributes of policy
type CreatePolicyResult struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}
