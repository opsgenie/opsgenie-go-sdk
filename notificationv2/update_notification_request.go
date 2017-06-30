package notificationv2

// UpdateNotificationRequest is a struct of request to update existing notification rule.
type UpdateNotificationRequest struct {
	*Identifier
	ApiKey           string
	Name             string             `json:"name"`
	Criteria         Criteria           `json:"criteria"`
	NotificationTime []NotificationTime `json:"notificationTime"`
	TimeRestriction  TimeRestriction    `json:"timeRestriction"`
	Schedules        []Schedule         `json:"schedules"`
	Steps            []Step             `json:"steps"`
	Repeat           Repeat             `json:"repeat"`
	Order            int                `json:"order"`
	Enabled          bool               `json:"enabled"`
}

// GetApiKey returns api key.
func (r *UpdateNotificationRequest) GetApiKey() string {
	return r.ApiKey
}
