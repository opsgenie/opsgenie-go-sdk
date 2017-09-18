package policyv1

const (
	ValueBasedDeduplicationActionType = "value-based"
	FrequencyBasedDeduplicationActionType = "frequency-based"
)

// DeduplicationActionType is a deduplication type. Should be one of "value-based" or "frequency-based"
type DeduplicationActionType string
