package schedule

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/syllabix/oncall/common/is"
	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/datastore/shift"
	"github.com/syllabix/oncall/datastore/user"
	"github.com/syllabix/oncall/service/oncall"
	"github.com/volatiletech/null/v8"
	"go.uber.org/zap"
)

var (
	ErrNotFound      = errors.New("schedule not found")
	ErrAlreadyExists = errors.New("there is already a schedule configured for this channel")
	ErrNoActiveShift = errors.New("there is no active shift for this schedule")
)

type Manager interface {
	Create(oncall.Schedule) (oncall.Schedule, error)
	GetAll() ([]oncall.Schedule, error)
	GetActiveShift(channelID string) (oncall.Shift, error)
	GetNextShifts(channelID string) ([]oncall.Shift, error)
	StartShift(scheduleID int) (oncall.Schedule, error)
	EndShift(scheduleID int) (oncall.Schedule, error)
}

func NewManager(db schedule.Store, users user.Store, shifts shift.Store, log *zap.Logger) Manager {
	return &manager{db, users, shifts, log}
}

type manager struct {
	db     schedule.Store
	users  user.Store
	shifts shift.Store

	log *zap.Logger
}

func (m *manager) Create(sched oncall.Schedule) (oncall.Schedule, error) {
	result, err := m.db.Create(context.TODO(), asModel(sched))
	if err != nil {
		if errors.Is(err, schedule.ErrAlreadyInUse) {
			return oncall.Schedule{}, ErrAlreadyExists
		}
		return oncall.Schedule{}, err
	}
	return asSchedule(result), nil
}

func (m *manager) GetAll() ([]oncall.Schedule, error) {
	results, err := m.db.GetEnabledSchedules(context.TODO())
	if err != nil {
		return nil, err
	}

	var schedules []oncall.Schedule
	for _, result := range results {
		schedules = append(schedules, asSchedule(result))
	}

	return schedules, nil
}

func (m *manager) GetNextShifts(channelID string) ([]oncall.Shift, error) {
	sched, err := m.db.GetByChannelID(context.Background(), channelID)
	if err != nil {
		if errors.Is(err, schedule.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to retrieve on call schedule for channel id: %w", err)
	}

	return calenderify(sched), nil
}

func (m *manager) GetActiveShift(channelID string) (oncall.Shift, error) {
	schedule, err := m.db.GetByChannelID(context.TODO(), channelID)
	if err != nil {
		return oncall.Shift{}, fmt.Errorf("failed to retrieve active on call shift: %w", err)
	}

	if isAfterHours(schedule) {
		return oncall.Shift{}, ErrNoActiveShift
	}

	if schedule.WeekdaysOnly && is.TheWeekend() {
		return oncall.Shift{}, ErrNoActiveShift
	}

	var user *model.User
loop:
	for _, shift := range schedule.R.Shifts {
		switch shift.Status.String {
		case model.ShiftStatusActive:
			user = shift.R.User
		case model.ShiftStatusOverride:
			user = shift.R.User
			break loop
		}
	}

	return oncall.Shift{
		UserID:      user.ID,
		SlackHandle: user.SlackHandle,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
	}, nil
}

func (m *manager) StartShift(scheduleID int) (oncall.Schedule, error) {
	schedule, err := m.db.GetByID(context.TODO(), scheduleID)
	if err != nil {
		return oncall.Schedule{},
			fmt.Errorf("failed to retrieve on call schedule: %w", err)
	}

	if len(schedule.R.Shifts) < 1 {
		return oncall.Schedule{}, ErrNoActiveShift
	}

	var updates []*model.Shift
	current, next := nextShiftFrom(schedule)
	if current != nil {
		current.Status = null.StringFromPtr(nil)
		current.StartedAt = null.TimeFromPtr(nil)
		updates = append(updates, current)
	}

	next.Status = null.StringFrom(model.ShiftStatusActive)
	next.StartedAt = null.TimeFrom(time.Now())
	updates = append(updates, next)

	err = m.shifts.Update(context.TODO(), updates...)
	if err != nil {
		return oncall.Schedule{},
			fmt.Errorf("failed to update on call schedule shifts: %w", err)
	}

	sched := asSchedule(schedule)
	sched.ActiveShift = &oncall.Shift{
		SlackHandle: next.R.User.SlackHandle,
	}

	return sched, nil
}

func (m *manager) EndShift(scheduleID int) (oncall.Schedule, error) {
	schedule, err := m.db.GetByID(context.TODO(), scheduleID)
	if err != nil {
		return oncall.Schedule{},
			fmt.Errorf("failed to retrieve on call schedule: %w", err)
	}

	if len(schedule.R.Shifts) < 1 {
		return oncall.Schedule{}, ErrNoActiveShift
	}

	override := findOverride(schedule)
	if override != nil {
		override.Status = null.StringFromPtr(nil)
		err = m.shifts.Update(context.TODO(), override)
		if err != nil {
			m.log.Error("failed to update override shift", zap.Error(err))
		}
	}

	current, _ := nextShiftFrom(schedule)

	sched := asSchedule(schedule)
	sched.ActiveShift = &oncall.Shift{
		SlackHandle: current.R.User.SlackHandle,
	}

	return sched, nil
}

func asModel(schedule oncall.Schedule) *model.Schedule {
	return &model.Schedule{
		Name:           schedule.Name,
		TeamSlackID:    schedule.TeamID,
		Interval:       string(schedule.Interval),
		WeekdaysOnly:   schedule.WeekdaysOnly,
		SlackChannelID: schedule.ChannelID,
		StartTime:      schedule.StartTime,
		EndTime:        schedule.EndTime,
	}
}

func asSchedule(model *model.Schedule) oncall.Schedule {
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
