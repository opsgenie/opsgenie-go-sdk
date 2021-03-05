package teamv2

const (
	// The list of action results.
	CreatedResult Result = "Created"
	DeletedResult Result = "Deleted"
)

// ActionResult contains result of action with notification rule.
type ActionResult struct {
	Result Result `json:"result"`
}

// Result contains string status of notification action result.
type Result string
