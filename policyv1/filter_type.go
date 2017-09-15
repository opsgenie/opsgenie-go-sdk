package policyv1

const (
	MatchAllConditionsFilterType FilterType = "match-all-conditions"
	MatchAllFilterType           FilterType = "match-all"
	MatchAnyCondition            FilterType = "match-any-condition"
)

type FilterType string
