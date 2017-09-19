package schedulev2

import "time"

type Rotation struct {
	Name         string        `json:"name"`
	StartDate    time.Time     `json:"startDate"`
	EndDate      time.Time     `json:"endDate"`
	Type         RotationType  `json:"type"`
	Length       int           `json:"length"`
	Participants []Participant `json:"participants"`
	TimeRestriction
}

type Participant struct {
	Type     ParticipantType `json:"type"`
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Username string          `json:"username"`
}

type TimeRestriction struct {
	Type         TimeRestrictionType `json:"type"`
	Restriction  Restriction         `json:"restriction"`
	Restrictions []Restriction       `json:"restrictions"`
}

type Restriction struct {
	StartDay  Day  `json:"startDay"`
	EndDay    Day  `json:"endDay"`
	StartHour *int `json:"startHour"`
	StartMin  *int `json:"startMin"`
	EndHour   *int `json:"endHour"`
	EndMin    *int `json:"endMin"`
}
