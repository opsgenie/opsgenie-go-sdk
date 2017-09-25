package contactv2

const (
	EmailContactMethod ContactMethod = "email"
	SmsContactMethod   ContactMethod = "sms"
	VoiceContactMethod ContactMethod = "voice"
)

type ContactMethod string
