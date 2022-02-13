package oncall

import "time"

type Interval string

const (
	Daily    Interval = "daily"
	Weekly   Interval = "weekly"
	BiWeekly Interval = "bi-weekly"
)

type Schedule struct {
	ID        int
	TeamID    string
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Interval  Interval
}
