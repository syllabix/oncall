package schedule

import (
	"errors"
	"fmt"

	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/datastore/user"
	"github.com/syllabix/oncall/service/schedule/oncall"
	"github.com/volatiletech/null/v8"
)

var (
	ErrAlreadyExists = errors.New("there is already a schedule configured for this channel")
)

type Manager interface {
	Create(oncall.Schedule) (oncall.Schedule, error)
	GetAll() ([]oncall.Schedule, error)
	GetNextShifts(channelID string) ([]oncall.Shift, error)
	StartShift(scheduleID string) (oncall.Schedule, error)
	EndShift(scheduleID string) (oncall.Schedule, error)
}

func NewManager(db schedule.Store, users user.Store) Manager {
	return &manager{db, users}
}

type manager struct {
	db    schedule.Store
	users user.Store
}

func (m *manager) Create(sched oncall.Schedule) (oncall.Schedule, error) {
	result, err := m.db.Create(asModel(sched))
	if err != nil {
		if errors.Is(err, schedule.ErrAlreadyInUse) {
			return oncall.Schedule{}, ErrAlreadyExists
		}
		return oncall.Schedule{}, err
	}
	return asSchedule(result), nil
}

func (m *manager) GetAll() ([]oncall.Schedule, error) {
	results, err := m.db.GetEnabledSchedules()
	if err != nil {
		return nil, err
	}

	var schedules []oncall.Schedule
	for _, result := range results {
		schedules = append(schedules, asSchedule(*result))
	}

	return schedules, nil
}

func (m *manager) GetNextShifts(channelID string) ([]oncall.Shift, error) {
	schedule, err := m.db.GetByChannelID(channelID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve on call schedule for channel id: %w", err)
	}

	userIDs := nextFiveShifts(schedule)
	users, err := m.users.GetAll(userIDs...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users for upcoming shifts: %w", err)
	}

	return asShifts(users, schedule), nil
}

func (m *manager) StartShift(scheduleID string) (oncall.Schedule, error) {
	schedule, err := m.db.GetByID(scheduleID)
	if err != nil {
		return oncall.Schedule{},
			fmt.Errorf("failed to retrieve on call schedule: %w", err)
	}

	next := nextShift(schedule.ActiveShift.String, schedule.Shifts)
	if len(next) < 1 {
		return asSchedule(schedule), nil
	}

	user, err := m.users.GetByID(next)
	if err != nil {
		return oncall.Schedule{}, fmt.Errorf("failed to get user for next shift: %w", err)
	}

	schedule.ActiveShift = null.StringFrom(next)
	_, err = m.db.Update(schedule)
	if err != nil {
		return oncall.Schedule{}, fmt.Errorf("failed to update schedule: %w", err)
	}

	sched := asSchedule(schedule)
	sched.ActiveShift = &oncall.UserOnDuty{
		SlackHandle: user.SlackHandle,
	}

	return sched, nil
}

func (m *manager) EndShift(scheduleID string) (oncall.Schedule, error) {
	schedule, err := m.db.GetByID(scheduleID)
	if err != nil {
		return oncall.Schedule{},
			fmt.Errorf("failed to retrieve on call schedule: %w", err)
	}

	if schedule.ActiveShift.IsZero() {
		return asSchedule(schedule), nil
	}

	user, err := m.users.GetByID(schedule.ActiveShift.String)
	if err != nil {
		return oncall.Schedule{}, fmt.Errorf("failed to get user for shift end: %w", err)
	}

	sched := asSchedule(schedule)
	sched.ActiveShift = &oncall.UserOnDuty{
		SlackHandle: user.SlackHandle,
	}

	return sched, nil
}

func asModel(schedule oncall.Schedule) model.Schedule {
	return model.Schedule{
		Name:           schedule.Name,
		TeamSlackID:    schedule.TeamID,
		Interval:       string(schedule.Interval),
		SlackChannelID: schedule.ChannelID,
		StartTime:      schedule.StartTime,
		EndTime:        schedule.EndTime,
	}
}

func asSchedule(model model.Schedule) oncall.Schedule {
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
