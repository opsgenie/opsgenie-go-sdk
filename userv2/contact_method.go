package userv2

const (
	// The list of notification methods.
	SMSNotifyMethod    = "sms"
	EmailNotifyMethod  = "email"
	VoiceNotifyMethod  = "voice"
	MobileNotifyMethod = "mobile"
)

// ContactMethod is a type of user contact method.
type ContactMethod string
