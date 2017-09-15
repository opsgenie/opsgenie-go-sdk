package policyv1

type Filter struct {
	Type       FilterType  `json:"type"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	Field         Field         `json:"field"`
	Operation     OperationType `json:"operation"`
	Key           string        `json:"key"`
	Not           bool          `json:"not"`
	ExpectedValue string        `json:"expectedValue"`
}
