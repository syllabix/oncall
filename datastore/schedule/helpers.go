package schedule

import (
	"github.com/syllabix/oncall/common/is"
	"github.com/syllabix/oncall/datastore/model"
)

func setActiveShift(schedule *model.Schedule) bool {
	if schedule.WeekdaysOnly && is.TheWeekend() {
		return false
	}

	for _, shift := range schedule.R.Shifts {
		if shift.Status.String == model.ShiftStatusActive {
			return false
		}
	}

	return true
}

func asShifts(scheduleID int, users model.UserSlice) (shifts model.ShiftSlice) {
	for _, user := range users {
		shifts = append(shifts, &model.Shift{
			UserID:     user.ID,
			ScheduleID: scheduleID,
		})
	}
	return shifts
}
