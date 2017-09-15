package policyv1

// EnablePolicyRequest is a request to enable policy by id
type EnablePolicyRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey return api key
func (r *EnablePolicyRequest) GetApiKey() string {
	return r.ApiKey
}

// EnablePolicyResponse is a response of enabling policy
type EnablePolicyResponse struct {
	ResponseMeta
	Status string `json:"result"`
}
