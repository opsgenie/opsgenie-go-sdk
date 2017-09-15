package policyv1

// DeletePolicyRequest is a request to delete policy by id
type DeletePolicyRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey return api key
func (r *DeletePolicyRequest) GetApiKey() string {
	return r.ApiKey
}

// DeletePolicyResponse is a response of deleting policy
type DeletePolicyResponse struct {
	ResponseMeta
	Status string `json:"result"`
}
