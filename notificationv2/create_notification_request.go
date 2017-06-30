package notificationv2

// CreateNotificationRequest is a struct of request to crate new notification rule.
type CreateNotificationRequest struct {
	*Identifier
	ApiKey           string
	Name             string             `json:"name"`
	ActionType       ActionType         `json:"actionType"`
	Criteria         Criteria           `json:"criteria"`
	NotificationTime []NotificationTime `json:"notificationTime"`
	TimeRestriction  *TimeRestriction   `json:"timeRestriction"`
	Schedules        []Schedule         `json:"schedules"`
	Order            int                `json:"order"`
	Steps            []Step             `json:"steps"`
	Repeat           Repeat             `json:"repeat"`
	Enabled          bool               `json:"enabled"`
}

// GetApiKey returns api key.
func (r *CreateNotificationRequest) GetApiKey() string {
	return r.ApiKey
}
