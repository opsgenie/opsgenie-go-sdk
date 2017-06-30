package notificationv2

// ListNotificationRequest is a struct of request to get list of existing notification rules.
type ListNotificationRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *ListNotificationRequest) GetApiKey() string {
	return r.ApiKey
}
