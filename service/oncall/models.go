package oncall

import "time"

type Interval = string

const (
	Daily    Interval = "daily"
	Weekly   Interval = "weekly"
	BiWeekly Interval = "bi-weekly"
	Monthy   Interval = "monthly"
)

type Shift struct {
	UserID       int
	FirstName    string
	LastName     string
	SlackHandle  string
	ScheduleName string
	StartTime    time.Time
	EndTime      time.Time
}

type Schedule struct {
	ID            int
	TeamID        string
	ChannelID     string
	Name          string
	StartTime     time.Time
	EndTime       time.Time
	Interval      Interval
	ActiveShift   *Shift
	OverrideShift *Shift
	WeekdaysOnly  bool
	IsEnabled     bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
