package shift

import "github.com/syllabix/oncall/datastore/model"

func findShift(userID int, schedule *model.Schedule) *model.Shift {
	if len(schedule.R.Shifts) < 1 {
		return nil
	}

	for _, shift := range schedule.R.Shifts {
		if shift.R.User.ID == userID {
			return shift
		}
	}

	return nil
}

func nextActive(schedule *model.Schedule) *model.Shift {
	if len(schedule.R.Shifts) < 1 {
		return nil
	}

	for i, shift := range schedule.R.Shifts {
		if shift.Status.String == model.ShiftStatusActive {
			if i == len(schedule.R.Shifts) {
				return schedule.R.Shifts[0]
			}
			return schedule.R.Shifts[i+1]
		}
	}

	return nil
}
