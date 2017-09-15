package policyv1

const (
	MatchesOperationType       OperationType = "matches"
	ContainsOpertionType       OperationType = "contains"
	StartsWithOperationType    OperationType = "starts-with"
	EndsWithOperationType      OperationType = "ends-with"
	EqualsOperationType        OperationType = "equals"
	ContainsKeyOperatonType    OperationType = "contains-key"
	ContainsValueOperationType OperationType = "contains-value"
	GreaterThanOperationType   OperationType = "greater-than"
	LessThanOperationType      OperationType = "less-than"
	IsEmptyOperationType       OperationType = "is-empty"
	EqualsIgnoreWhitespace     OperationType = "equals-ignore-whitespace"
)

type OperationType string
