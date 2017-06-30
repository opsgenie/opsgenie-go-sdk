package notificationv2

const (
	// List of time periods that notification for schedule start/end will be sent.
	JustBeforeNotificationTime        = "just-before"
	FifteenMinutesAgoNotificationTime = "15-minutes-ago"
	OneHourAgoNotificationTime        = "1-hour-ago"
	OneDayAgoNotificationTime         = "1-day-ago"
)

// NotificationTime is type of time periods that notification for start/end will be sent.
type NotificationTime string
