package schedule

import (
	"time"

	"github.com/syllabix/oncall/datastore/entity"
	"github.com/syllabix/oncall/datastore/entity/schedule"
	"github.com/syllabix/oncall/datastore/entity/shift"
	"github.com/syllabix/oncall/service/oncall"
)

func asModel(s oncall.Schedule) *entity.Schedule {
	return &entity.Schedule{
		Name:           s.Name,
		TeamSlackID:    s.TeamID,
		Interval:       schedule.Interval(s.Interval),
		WeekdaysOnly:   s.WeekdaysOnly,
		SlackChannelID: s.ChannelID,
		StartTime:      s.StartTime,
		EndTime:        s.EndTime,
	}
}

func asSchedule(model *entity.Schedule) oncall.Schedule {
	return oncall.Schedule{
		ID:           model.ID,
		Name:         model.Name,
		TeamID:       model.TeamSlackID,
		Interval:     oncall.Interval(model.Interval),
		ChannelID:    model.SlackChannelID,
		WeekdaysOnly: model.WeekdaysOnly,
		IsEnabled:    model.IsEnabled,
		StartTime:    model.StartTime,
		EndTime:      model.EndTime,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}

func findOverride(schedule *entity.Schedule) *entity.Shift {
	for _, s := range schedule.Edges.Shifts {
		if s.Status == nil {
			continue
		}

		if *s.Status == shift.StatusOverride {
			return s
		}
	}
	return nil
}

func nextShiftFrom(schedule *entity.Schedule) (current *entity.Shift, next *entity.Shift) {
	idx := 0
	for i, s := range schedule.Edges.Shifts {
		if s.Status == nil {
			continue
		}

		if *s.Status == shift.StatusActive {
			idx = i + 1
			current = s
			break
		}
	}

	numShifts := len(schedule.Edges.Shifts)
	if idx == numShifts {
		return current, schedule.Edges.Shifts[0]
	}

	return current, schedule.Edges.Shifts[idx]
}

func isAfterHours(schedule *entity.Schedule) bool {
	now := time.Now()
	start := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		schedule.StartTime.Hour(),
		schedule.StartTime.Minute(),
		schedule.StartTime.Second(),
		schedule.StartTime.Nanosecond(),
		time.UTC,
	)
	end := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		schedule.EndTime.Hour(),
		schedule.EndTime.Minute(),
		schedule.EndTime.Second(),
		schedule.EndTime.Nanosecond(),
		time.UTC,
	)
	if now.Before(start) || now.After(end) {
		return true
	}
	return false
}

func arrange(shifts []*entity.Shift) (active *entity.Shift, ordered []*entity.Shift) {
	if len(shifts) < 1 {
		return nil, nil
	}

	idx := -1
	for i, s := range shifts {
		if s.Status == nil {
			continue
		}
		if *s.Status == shift.StatusActive {
			idx = i + 1
			active = s
			break
		}
	}

	if idx == -1 {
		// there are no active schedules yet
		// let's set the started at time to now
		// so when we compute the calendar we have a valid
		// date to start from
		offset := time.Now().AddDate(0, 0, -1)
		shifts[0].StartedAt = &offset
		return shifts[0], shifts
	}

	numShifts := len(shifts)
	ordered = make([]*entity.Shift, 0, numShifts)
	for len(ordered) != numShifts {
		if idx > len(shifts)-1 {
			idx = 0
		}
		ordered = append(ordered, shifts[idx])
		idx++
	}

	return active, ordered
}

func calenderify(sched *entity.Schedule) (shifts []oncall.Shift) {
	active, ordered := arrange(sched.Edges.Shifts)
	if len(ordered) < 1 {
		return shifts
	}

	dayIncr := 0
	monthIncr := 0

	switch sched.Interval {
	case schedule.IntervalDaily:
		dayIncr = 1
	case schedule.IntervalWeekly:
		dayIncr = 7
	case schedule.IntervalBiWeekly:
		dayIncr = 14
	case schedule.IntervalMonthly:
		monthIncr = 1
	}

	start := active.StartedAt.AddDate(0, monthIncr, dayIncr)
	for _, shift := range ordered {
		if sched.WeekdaysOnly {
			switch start.Weekday() {
			case time.Saturday:
				start = start.AddDate(0, 0, 2)
			case time.Sunday:
				start = start.AddDate(0, 0, 1)
			}
		}
		user := shift.Edges.User
		shifts = append(shifts, oncall.Shift{
			UserID:      user.ID,
			SlackHandle: user.SlackHandle,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			StartTime:   time.Date(start.Year(), start.Month(), start.Day(), sched.StartTime.Hour(), 0, 0, 0, time.UTC),
			EndTime:     time.Date(start.Year(), start.Month(), start.Day(), sched.EndTime.Hour(), 0, 0, 0, time.UTC),
		})

		start = start.AddDate(0, monthIncr, dayIncr)
	}

	return shifts
}
