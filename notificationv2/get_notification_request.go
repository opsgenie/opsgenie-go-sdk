package notificationv2

// GetNotificationRequest is a struct of request to get notification rule.
type GetNotificationRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *GetNotificationRequest) GetApiKey() string {
	return r.ApiKey
}
