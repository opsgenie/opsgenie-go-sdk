package notificationv2

// EnableNotificationRequest is a struct of request to enable specified notification rule.
type EnableNotificationRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *EnableNotificationRequest) GetApiKey() string {
	return r.ApiKey
}
