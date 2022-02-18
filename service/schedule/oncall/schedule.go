package oncall

import "time"

type Interval string

const (
	Daily    Interval = "daily"
	Weekly   Interval = "weekly"
	BiWeekly Interval = "bi-weekly"
)

type UserOnDuty struct {
	SlackHandle string
}

type Shift struct {
	UserID      string
	FirstName   string
	LastName    string
	SlackHandle string
	StartTime   time.Time
	EndTime     time.Time
}

type Schedule struct {
	ID            string
	TeamID        string
	ChannelID     string
	Name          string
	StartTime     time.Time
	EndTime       time.Time
	Interval      Interval
	ActiveShift   *UserOnDuty
	OverrideShift *UserOnDuty
	WeekdaysOnly  bool
	IsEnabled     bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
