package schedulev2

const (
	DailyRotationType  RotationType = "daily"
	WeeklyRotationType RotationType = "weekly"
	HourlyRotationType RotationType = "hourly"
)

type RotationType string
