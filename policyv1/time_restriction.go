package policyv1

type TimeRestrictions struct {
	Type         TimeRestrictionType `json:"type"`
	Restriction  Restriction         `json:"restriction"`
	Restrictions []Restriction       `json:"restrictions"`
}

type Restriction struct {
	StartDay  Day  `json:"startDay,omitempty"`
	EndDay    Day  `json:"endDay,omitempty"`
	StartHour *int `json:"startHour,omitempty"`
	EndHour   *int `json:"endHour,omitempty"`
	StartMin  *int `json:"startMin,omitempty"`
	EndMin    *int `json:"endMin,omitempty"`
}
