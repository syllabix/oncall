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

type SwapRequest struct {
	ChannelID    string
	UserSlackID1 string
	UserSlackID2 string
}

type Manager interface {
	SwapShift(context.Context, SwapRequest) (next *oncall.Shift, err error)
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

func (m *manager) SwapShift(ctx context.Context, req SwapRequest) (next *oncall.Shift, err error) {
	users, err := m.users.ListBySlackID(req.UserSlackID1, req.UserSlackID2)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users to swap schedule: %w", err)
	}

	if len(users) != 2 {
		return nil, fmt.Errorf("could not find both users requested for shift swap")
	}

	schedule, err := m.schedules.GetByChannelID(ctx, req.ChannelID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch schedule while removing shift: %w", err)
	}

	swapped, err := swapShifts(users[0], users[1], schedule)
	if err != nil {
		return nil, err
	}

	err = m.shifts.Update(ctx, swapped...)
	if err != nil {
		return nil, err
	}

	for _, shift := range swapped {
		if shift.Status.String == model.ShiftStatusActive {
			return asShift(shift, schedule), nil
		}
	}

	return nil, nil
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

	return asShift(newActive, schedule), nil
}
