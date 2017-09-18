package policyv1

const (
	UserRecipient       = "user"
	EscalationRecipient = "escalation"
	TeamRecipient       = "team"
	ScheduleRecipient   = "schedule"
	NoneRecipient       = "none"
	AllRecipient        = "all"
)

// Recipient is a list of recipients. It can be a list of escalations, schedules, teams, users or the reserved word
// none or all
type Recipient struct {
	Type     string `json:"type"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
