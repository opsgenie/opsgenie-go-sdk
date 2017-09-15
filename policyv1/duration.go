package policyv1

type Duration struct {
	TimeAmount *int     `json:"timeAmount,omitempty"`
	TimeUnit   TimeUnit `json:"timeUnit,omitempty"`
}
