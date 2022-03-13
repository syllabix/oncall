package schedule

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/syllabix/oncall/common/is"
	"github.com/syllabix/oncall/datastore/entity"
	"github.com/syllabix/oncall/datastore/entity/shift"
	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/datastore/schedule"
	shifts "github.com/syllabix/oncall/datastore/shift"
	"github.com/syllabix/oncall/datastore/user"
	"github.com/syllabix/oncall/service/oncall"
	"go.uber.org/zap"
)

var (
	ErrNotFound      = errors.New("schedule not found")
	ErrAlreadyExists = errors.New("there is already a schedule configured for this channel")
	ErrNoActiveShift = errors.New("there is no active shift for this schedule")
)

type Manager interface {
	Create(context.Context, oncall.Schedule) (oncall.Schedule, error)
	GetAll() ([]oncall.Schedule, error)
	GetActiveShift(channelID string) (oncall.Shift, error)
	GetNextShifts(channelID string) ([]oncall.Shift, error)
	StartShift(scheduleID int) (oncall.Schedule, error)
	EndShift(scheduleID int) (oncall.Schedule, error)
}

func NewManager(db schedule.Store, users user.Store, shifts shifts.Store, log *zap.Logger) Manager {
	return &manager{db, users, shifts, log}
}

type manager struct {
	db     schedule.Store
	users  user.Store
	shifts shifts.Store

	log *zap.Logger
}

func (m *manager) Create(ctx context.Context, sched oncall.Schedule) (oncall.Schedule, error) {
	result, err := m.db.Create(ctx, asModel(sched))
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

	var user *entity.User
loop:
	for _, s := range schedule.Edges.Shifts {
		if s.Status == nil {
			continue
		}

		switch *s.Status {
		case shift.StatusActive:
			user = s.Edges.User
		case model.ShiftStatusOverride:
			user = s.Edges.User
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

	if len(schedule.Edges.Shifts) < 1 {
		return oncall.Schedule{}, ErrNoActiveShift
	}

	var updates []*entity.Shift
	current, next := nextShiftFrom(schedule)
	if current != nil {
		current.Status = nil
		current.StartedAt = nil
		updates = append(updates, current)
	}

	status, now := shift.StatusActive, time.Now()
	next.Status = &status
	next.StartedAt = &now
	updates = append(updates, next)

	err = m.shifts.Update(context.TODO(), updates...)
	if err != nil {
		return oncall.Schedule{},
			fmt.Errorf("failed to update on call schedule shifts: %w", err)
	}

	sched := asSchedule(schedule)
	sched.ActiveShift = &oncall.Shift{
		SlackHandle: next.Edges.User.SlackHandle,
	}

	return sched, nil
}

func (m *manager) EndShift(scheduleID int) (oncall.Schedule, error) {
	schedule, err := m.db.GetByID(context.TODO(), scheduleID)
	if err != nil {
		return oncall.Schedule{},
			fmt.Errorf("failed to retrieve on call schedule: %w", err)
	}

	if len(schedule.Edges.Shifts) < 1 {
		return oncall.Schedule{}, ErrNoActiveShift
	}

	override := findOverride(schedule)
	if override != nil {
		override.Status = nil
		err = m.shifts.Update(context.TODO(), override)
		if err != nil {
			m.log.Error("failed to update override shift", zap.Error(err))
		}
	}

	current, _ := nextShiftFrom(schedule)

	sched := asSchedule(schedule)
	sched.ActiveShift = &oncall.Shift{
		SlackHandle: current.Edges.User.SlackHandle,
	}

	return sched, nil
}
