package notificationv2

const (
	// The list of action results.
	EnabledResult = "Enabled"
	DisableResult = "Disabled"
	DeletedResult = "Deleted"
)

// ActionResult contains result of action with notification rule.
type ActionResult struct {
	Result Result `json:"result"`
}

// Result contains string status of notification action result.
type Result string
