package policyv1

const (
	MatchAllConditionsFilterType FilterType = "match-all-conditions"
	MatchAllFilterTypeFilterType FilterType = "match-all"
	MatchAnyConditionFilterType  FilterType = "match-any-condition"
)

type FilterType string
