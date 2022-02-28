package schedule

import (
	"time"

	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/service/oncall"
)

func nextShiftFrom(schedule *model.Schedule) (current *model.Shift, next *model.Shift) {
	idx := 0
	for i, shift := range schedule.R.Shifts {
		if shift.Status.String == model.ShiftStatusActive {
			idx = i + 1
			current = shift
			break
		}
	}

	numShifts := len(schedule.R.Shifts)
	if idx == numShifts {
		return current, schedule.R.Shifts[0]
	}

	return current, schedule.R.Shifts[idx]
}

func isAfterHours(schedule *model.Schedule) bool {
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

func arrange(shifts model.ShiftSlice) (active *model.Shift, ordered model.ShiftSlice) {
	if len(shifts) < 1 {
		return nil, nil
	}

	idx := -1
	for i, shift := range shifts {
		if shift.Status.String == model.ShiftStatusActive {
			idx = i + 1
			active = shift
			break
		}
	}

	// there are no active schedules yet
	// let's set the started at time to now
	// so when we compute the calendar we have a valid
	// date to start from
	if idx == -1 {
		shifts[0].StartedAt.Time = time.Now().AddDate(0, 0, -1)
		return shifts[0], shifts
	}

	numShifts := len(shifts)
	ordered = make(model.ShiftSlice, 0, numShifts)
	for len(ordered) != numShifts {
		if idx > len(shifts)-1 {
			idx = 0
		}
		ordered = append(ordered, shifts[idx])
		idx++
	}

	return active, ordered
}

func calenderify(schedule *model.Schedule) (shifts []oncall.Shift) {
	active, ordered := arrange(schedule.R.Shifts)
	if len(ordered) < 1 {
		return shifts
	}

	dayIncr := 0
	monthIncr := 0

	switch schedule.Interval {
	case oncall.Daily:
		dayIncr = 1
	case oncall.Weekly:
		dayIncr = 7
	case oncall.BiWeekly:
		dayIncr = 14
	case oncall.Monthy:
		monthIncr = 1
	}

	start := active.StartedAt.Time.AddDate(0, monthIncr, dayIncr)
	for _, shift := range ordered {
		if schedule.WeekdaysOnly {
			switch start.Weekday() {
			case time.Saturday:
				start = start.AddDate(0, 0, 2)
			case time.Sunday:
				start = start.AddDate(0, 0, 1)
			}
		}
		user := shift.R.User
		shifts = append(shifts, oncall.Shift{
			UserID:      user.ID,
			SlackHandle: user.SlackHandle,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			StartTime:   time.Date(start.Year(), start.Month(), start.Day(), schedule.StartTime.Hour(), 0, 0, 0, time.UTC),
			EndTime:     time.Date(start.Year(), start.Month(), start.Day(), schedule.EndTime.Hour(), 0, 0, 0, time.UTC),
		})

		start = start.AddDate(0, monthIncr, dayIncr)
	}

	return shifts
}
