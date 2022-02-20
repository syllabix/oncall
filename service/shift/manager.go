package shift

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/datastore/shift"
	"github.com/syllabix/oncall/datastore/user"
	"github.com/syllabix/oncall/service/oncall"
	"github.com/volatiletech/null/v8"
)

var (
	ErrNotFound = errors.New("the shift does not exist")
)

type ErrScheduleEmpty struct {
	ID   int
	Name string
}

func (e *ErrScheduleEmpty) Error() string {
	return "this schedule no longer has any shifts assigned"
}

type RemoveRequest struct {
	ChannelID   string
	UserSlackID string
}

type Manager interface {
	RemoveShift(context.Context, RemoveRequest) (next *oncall.Shift, err error)
}

func NewManager(
	users user.Store,
	shifts shift.Store,
	schedules schedule.Store,
) Manager {
	return &manager{users, shifts, schedules}
}

type manager struct {
	users     user.Store
	shifts    shift.Store
	schedules schedule.Store
}

func (m *manager) RemoveShift(ctx context.Context, req RemoveRequest) (*oncall.Shift, error) {
	user, err := m.users.GetBySlackID(ctx, req.UserSlackID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user to remove from schedule: %w", err)
	}

	schedule, err := m.schedules.GetByChannelID(ctx, req.ChannelID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch schedule while removing shift: %w", err)
	}

	shiftToRemove := findShift(user.ID, schedule)
	if shiftToRemove == nil {
		return nil, ErrNotFound
	}

	err = m.shifts.Delete(ctx, shiftToRemove)
	if err != nil {
		return nil, fmt.Errorf("failed to delete shift: %w", err)
	}

	if len(schedule.R.Shifts) <= 1 {
		return nil, &ErrScheduleEmpty{
			ID:   schedule.ID,
			Name: schedule.Name,
		}
	}

	var newActive *model.Shift
	if shiftToRemove.Status.String == model.ShiftStatusActive {
		newActive = nextActive(schedule)
		if newActive != nil && newActive != shiftToRemove {
			newActive.Status = null.StringFrom(model.ShiftStatusActive)
			newActive.StartedAt = null.TimeFrom(time.Now())
		}
	}

	if newActive == nil {
		return nil, nil
	}

	err = m.shifts.Update(ctx, newActive)
	if err != nil {
		return nil, err
	}

	usr := newActive.R.User
	return &oncall.Shift{
		UserID:       usr.ID,
		FirstName:    usr.FirstName,
		LastName:     usr.LastName,
		SlackHandle:  usr.SlackHandle,
		ScheduleName: schedule.Name,
		StartTime:    schedule.StartTime,
		EndTime:      schedule.EndTime,
	}, nil
}
