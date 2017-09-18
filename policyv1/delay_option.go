package policyv1

const (
	ForDurationDelayOption   = "for-duration"
	NextTimeDelayOption      = "next-time"
	NextWeekdayDelayOption   = "next-weekday"
	NextMondayDelayOption    = "next-monday"
	NextTuesdayDelayOption   = "next-tuesday"
	NextWednesdayDelayOption = "next-wednesday"
	NextThursdayDelayOption  = "next-thursday"
	NextFridayDelayOption    = "next-friday"
	NextSaturdayDelayOption  = "next-saturday"
	NextSundayDelayOption    = "next-sunday"
)

// DelayOption is a delay type, which should be one of "for-duration", "next-time", "next-weekday", "next-monday",
// "next-tuesday", "next-wednesday", "next-thursday", "next-friday", "next-saturday", "next-sunday"
type DelayOption string
