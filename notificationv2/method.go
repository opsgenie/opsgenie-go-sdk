package notificationv2

const (
	// The list of notification method.
	SMSNotifyMethod    = "sms"
	EmailNotifyMethod  = "email"
	VoiceNotifyMethod  = "voice"
	MobileNotifyMethod = "mobile"
)

// Method is a method of notification.
type Method string
