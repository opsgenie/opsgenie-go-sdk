package notificationv2

const (
	// The list of condition operation, which are used for build notification criteria.
	MatchesConditionOperation = "matches"
	EqualsConditionOperation  = "equals"
	IsEmptyConditionOperation = "is-empty"
)

// Operation is type of matching operation, which is used to build filter of alerts.
type Operation string
