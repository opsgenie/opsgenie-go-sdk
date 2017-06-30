package notificationv2

const (
	// The list of statuses, which are used for enable/disable notification rules.
	EnableStatusAction  = "enable"
	DisableStatusAction = "disable"
)

// StatusAction is a status of notification action, which should be "enable" or "disable".
type StatusAction string
