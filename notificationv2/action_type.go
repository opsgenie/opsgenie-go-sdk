package notificationv2

const (
	// Types of action that notification rule will have after creating. It is used to create an alert.
	CreateAlertActionType         = "create-alert"
	AcknowledgedAlertActionType   = "acknowledged-alert"
	ClosedAlertActionType         = "closed-alert"
	AssignedAlertActionType       = "assigned-alert"
	AddNoteActionType             = "add-note"
	ScheduleStartActionType       = "schedule-start"
	ScheduleEndActionType         = "schedule-end"
	IncomingCallRoutingActionType = "incoming-call-routing"
)

// ActionType is the type of notification action. Instead of
type ActionType string
