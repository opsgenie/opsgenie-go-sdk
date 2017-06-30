package userv2

const (
	// EscalationsEntity is the query parameter to load user escalations.
	EscalationsEntity = "escalations"
	// EscalationsEntity is the query parameter to load user teams.
	TeamsEntity = "teams"
	// EscalationsEntity is the query parameter to load user forwarding rules.
	ForwardingRulesEntity = "forwarding-rules"
	// EscalationsEntity is the query parameter to load user schedules.
	SchedulesEntity = "schedules"
)

// Entity is a name of user entity, for example: escalation, team, forwarding-rule, schedule.
type Entity string
