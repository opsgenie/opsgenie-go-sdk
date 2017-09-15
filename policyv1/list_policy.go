package policyv1

// ListPolicyRequest is a request to load all policies
type ListPolicyRequest struct {
	*Identifier
	ApiKey            string
}

// GetApiKey returns api key
func (r *ListPolicyRequest) GetApiKey() string {
	return r.ApiKey
}

// ListPolicyResponse is a response of loading all policies
type ListPolicyResponse struct {
	ResponseMeta
	Policies []ListPolicyResult `json:"data"`
}

// ListPolicyResult contains data of policy
type ListPolicyResult struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Order   int    `json:"order"`
	Enabled bool   `json:"enabled"`
}
