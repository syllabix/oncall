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

func nextShifts(count int, schedule model.Schedule) (shifts []string) {

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

	for len(shifts) < count {
		if idx > len(schedule.Shifts)-1 {
			idx = 0
		}
		shifts = append(shifts, schedule.Shifts[idx])
		idx++
	}

	return shifts
}

func asShifts(users []model.User, orderUserIDs []string, schedule model.Schedule) (shifts []oncall.Shift) {
	now := time.Now()
	day := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24)

	userMap := make(map[string]model.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	for _, userID := range orderUserIDs {
		if schedule.WeekdaysOnly {
			switch day.Weekday() {
			case time.Saturday:
				day = day.Add(time.Hour * 48)
			case time.Sunday:
				day = day.Add(time.Hour * 24)
			}
		}

		user := userMap[userID]

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
