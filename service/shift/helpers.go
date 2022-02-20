package shift

import (
	"errors"
	"fmt"

	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/service/oncall"
)

func asShift(s *model.Shift, schedule *model.Schedule) *oncall.Shift {
	usr := s.R.User
	return &oncall.Shift{
		UserID:       usr.ID,
		FirstName:    usr.FirstName,
		LastName:     usr.LastName,
		SlackHandle:  usr.SlackHandle,
		ScheduleName: schedule.Name,
		StartTime:    schedule.StartTime,
		EndTime:      schedule.EndTime,
	}
}

func extractShift(userID int, schedule *model.Schedule) *model.Shift {
	if len(schedule.R.Shifts) < 1 {
		return nil
	}

	for i, shift := range schedule.R.Shifts {
		if shift.R.User.ID == userID {
			schedule.R.Shifts = append(schedule.R.Shifts[:i], schedule.R.Shifts[i+1:]...)
			return shift
		}
	}

	return nil
}

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
			if i == len(schedule.R.Shifts)-1 {
				return schedule.R.Shifts[0]
			}
			return schedule.R.Shifts[i+1]
		}
	}

	return nil
}

func currentOverride(schedule *model.Schedule) *model.Shift {
	if len(schedule.R.Shifts) < 1 {
		return nil
	}

	for _, shift := range schedule.R.Shifts {
		if shift.Status.String == model.ShiftStatusOverride {
			return shift
		}
	}

	return nil
}

func swapShifts(userA, userB model.User, schedule *model.Schedule) ([]*model.Shift, error) {
	if len(schedule.R.Shifts) < 2 {
		return nil, errors.New("the schedule has less then two shifts")
	}

	set := make(map[int]*model.Shift)
	for _, shift := range schedule.R.Shifts {
		set[shift.UserID] = shift
	}

	shiftA, contains := set[userA.ID]
	if !contains {
		return nil, fmt.Errorf("user %d is not in schedule", userA.ID)
	}

	shiftB, contains := set[userB.ID]
	if !contains {
		return nil, fmt.Errorf("user %d is not in schedule", userB.ID)
	}

	return []*model.Shift{
		{
			UserID:     shiftA.UserID,
			ScheduleID: shiftA.ScheduleID,
			R:          shiftA.R,
			Status:     shiftB.Status,
			StartedAt:  shiftB.StartedAt,
			SequenceID: shiftB.SequenceID,
		},
		{
			UserID:     shiftB.UserID,
			ScheduleID: shiftB.ScheduleID,
			R:          shiftB.R,
			Status:     shiftA.Status,
			StartedAt:  shiftA.StartedAt,
			SequenceID: shiftA.SequenceID,
		},
	}, nil

}
