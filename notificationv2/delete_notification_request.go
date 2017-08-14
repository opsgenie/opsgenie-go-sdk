package notificationv2

// DeleteNotificationRequest is a struct of request to delete existing notification rule.
type DeleteNotificationRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *DeleteNotificationRequest) GetApiKey() string {
	return r.ApiKey
}
