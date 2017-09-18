package policyv1

const (
	P1Priority = "P1"
	P2Priority = "P2"
	P3Priority = "P3"
	P4Priority = "P4"
	P5Priority = "P5"
)

// Priority is a priority of an alert. Should be one of "P1", "P2", "P3", "P4", or "P5"
type Priority string
