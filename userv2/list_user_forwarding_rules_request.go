package userv2

// ListUserForwardingRulesRequest is a request for getting list of forwarding rules.
type ListUserForwardingRulesRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListUserForwardingRulesRequest) GetApiKey() string {
	return r.ApiKey
}
