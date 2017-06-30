package notificationv2

// DisableNotificationRequest is a struct of request to disable specified notification rule.
type DisableNotificationRequest struct {
	*Identifier
	ApiKey string
}

// GetApiKey returns api key.
func (r *DisableNotificationRequest) GetApiKey() string {
	return r.ApiKey
}
