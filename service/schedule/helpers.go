package schedule

import (
	"time"

	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/service/schedule/oncall"
)

func nextShift(activeShift string, shifts []string) string {
	if len(shifts) < 1 {
		return ""
	}

	if activeShift == "" {
		return shifts[0]
	}

	for i, shift := range shifts {
		if shift == activeShift {
			if len(shifts) == i+1 {
				return shifts[0]
			}
			return shifts[i+1]
		}
	}

	return ""
}

func nextFiveShifts(schedule model.Schedule) (shifts []string) {

	var idx int

	if schedule.ActiveShift.IsZero() {
		idx = 0
	} else {
		for i, id := range schedule.Shifts {
			if id == schedule.ActiveShift.String {
				idx = i + 1
			}
		}
	}

	for len(shifts) < 5 {
		if idx > len(schedule.Shifts)-1 {
			idx = 0
		}
		shifts = append(shifts, schedule.Shifts[idx])
		idx++
	}

	return shifts
}

func asShifts(users []model.User, schedule model.Schedule) (shifts []oncall.Shift) {
	day := time.Now().Add(time.Hour * 24)

	for _, user := range users {
		if schedule.WeekdaysOnly {
			switch day.Weekday() {
			case time.Saturday:
				day = day.Add(time.Hour * 48)
			case time.Sunday:
				day = day.Add(time.Hour * 24)
			}
		}

		shifts = append(shifts, oncall.Shift{
			UserID:      user.ID,
			SlackHandle: user.SlackHandle,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			StartTime:   time.Date(day.Year(), day.Month(), day.Day(), schedule.StartTime.Hour(), 0, 0, 0, time.UTC),
			EndTime:     time.Date(day.Year(), day.Month(), day.Day(), schedule.EndTime.Hour(), 0, 0, 0, time.UTC),
		})
		day = day.Add(time.Hour * 24)
	}

	return shifts
}
