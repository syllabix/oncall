package shift

import (
	"errors"
	"fmt"

	"github.com/syllabix/oncall/datastore/entity"
	"github.com/syllabix/oncall/datastore/entity/shift"
	"github.com/syllabix/oncall/service/oncall"
)

func asShift(s *entity.Shift, schedule *entity.Schedule) *oncall.Shift {
	usr := s.Edges.User
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

func extractShift(userID int, schedule *entity.Schedule) *entity.Shift {
	if len(schedule.Edges.Shifts) < 1 {
		return nil
	}

	for i, shift := range schedule.Edges.Shifts {
		if shift.Edges.User.ID == userID {
			schedule.Edges.Shifts = append(schedule.Edges.Shifts[:i], schedule.Edges.Shifts[i+1:]...)
			return shift
		}
	}

	return nil
}

func findShift(userID int, schedule *entity.Schedule) *entity.Shift {
	if len(schedule.Edges.Shifts) < 1 {
		return nil
	}

	for _, shift := range schedule.Edges.Shifts {
		if shift.Edges.User.ID == userID {
			return shift
		}
	}

	return nil
}

func nextActive(schedule *entity.Schedule) *entity.Shift {
	if len(schedule.Edges.Shifts) < 1 {
		return nil
	}

	for i, s := range schedule.Edges.Shifts {
		if s.Status == nil {
			continue
		}

		if *s.Status == shift.StatusActive {
			if i == len(schedule.Edges.Shifts)-1 {
				return schedule.Edges.Shifts[0]
			}
			return schedule.Edges.Shifts[i+1]
		}
	}

	return nil
}

func currentOverride(schedule *entity.Schedule) *entity.Shift {
	if len(schedule.Edges.Shifts) < 1 {
		return nil
	}

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

func swapShifts(userA, userB *entity.User, schedule *entity.Schedule) ([]*entity.Shift, error) {
	if len(schedule.Edges.Shifts) < 2 {
		return nil, errors.New("the schedule has less then two shifts")
	}

	set := make(map[int]*entity.Shift)
	for _, shift := range schedule.Edges.Shifts {
		set[shift.Edges.User.ID] = shift
	}

	shiftA, contains := set[userA.ID]
	if !contains {
		return nil, fmt.Errorf("user %d is not in schedule", userA.ID)
	}

	shiftB, contains := set[userB.ID]
	if !contains {
		return nil, fmt.Errorf("user %d is not in schedule", userB.ID)
	}

	return []*entity.Shift{
		{
			Status:     shiftB.Status,
			StartedAt:  shiftB.StartedAt,
			SequenceID: shiftB.SequenceID,
			Edges: entity.ShiftEdges{
				User:     shiftA.Edges.User,
				Schedule: shiftA.Edges.Schedule,
			},
		},
		{
			Status:     shiftA.Status,
			StartedAt:  shiftA.StartedAt,
			SequenceID: shiftA.SequenceID,
			Edges: entity.ShiftEdges{
				User:     shiftB.Edges.User,
				Schedule: shiftB.Edges.Schedule,
			},
		},
	}, nil

}
