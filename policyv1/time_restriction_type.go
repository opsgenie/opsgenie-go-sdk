package policyv1

const (
	TimeOfDayTimeRestrictionType           TimeRestrictionType = "time-of-day"
	WeekdayAndTimeOfDayTimeRestrictionType TimeRestrictionType = "weekday-and-time-of-day"
)

type TimeRestrictionType string
