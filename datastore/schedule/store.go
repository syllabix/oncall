package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/syllabix/oncall/datastore/entity"
	"github.com/syllabix/oncall/datastore/entity/schedule"
	"github.com/syllabix/oncall/datastore/entity/shift"
)

type Store struct {
	db     *entity.ScheduleClient
	client *entity.Client
}

func NewStore(db *entity.Client) Store {
	return Store{db.Schedule, db}
}

func (s *Store) Create(ctx context.Context, schedule *entity.Schedule) (*entity.Schedule, error) {
	schedule, err := s.db.Create().
		SetSlackChannelID(schedule.SlackChannelID).
		SetTeamSlackID(schedule.TeamSlackID).
		SetName(schedule.Name).
		SetInterval(schedule.Interval).
		SetIsEnabled(schedule.IsEnabled).
		SetEndTime(schedule.EndTime).
		SetStartTime(schedule.StartTime).
		SetWeekdaysOnly(schedule.WeekdaysOnly).
		Save(ctx)
	if err != nil {
		return failure[*entity.Schedule](
			fmt.Errorf("could not create schedule: %w", err),
		)
	}

	return schedule, nil
}

func (s *Store) Update(ctx context.Context, schedule *entity.Schedule) (*entity.Schedule, error) {
	schedule, err := s.db.UpdateOne(schedule).Save(ctx)
	if err != nil {
		return failure[*entity.Schedule](
			fmt.Errorf("could not update schedule: %w", err),
		)
	}
	return schedule, nil
}

func (s *Store) GetByID(ctx context.Context, id int) (*entity.Schedule, error) {
	schedule, err := s.db.Query().
		Where(schedule.IDEQ(id)).
		WithShifts(func(q *entity.ShiftQuery) {
			q.Order(entity.Asc(shift.FieldSequenceID))
			q.WithUser()
		}).
		Only(ctx)
	if err != nil {
		return failure[*entity.Schedule](
			fmt.Errorf("could not get schedule by id: %w", err),
		)
	}
	return schedule, nil
}

func (s *Store) GetByChannelID(ctx context.Context, id string) (*entity.Schedule, error) {
	schedule, err := s.db.Query().
		Where(schedule.SlackChannelIDEQ(id)).
		WithShifts(func(q *entity.ShiftQuery) {
			q.Order(entity.Asc(shift.FieldSequenceID))
			q.WithUser()
		}).
		Only(ctx)
	if err != nil {
		return failure[*entity.Schedule](
			fmt.Errorf("could not get schedule by channel id: %w", err),
		)
	}
	return schedule, nil
}

func (s *Store) GetEnabledSchedules(ctx context.Context) ([]*entity.Schedule, error) {
	schedules, err := s.db.Query().
		Where(schedule.IsEnabledEQ(true)).
		WithShifts(func(q *entity.ShiftQuery) {
			q.Order(entity.Asc(shift.FieldSequenceID))
			q.WithUser()
		}).
		All(ctx)
	if err != nil {
		return failure[[]*entity.Schedule](
			fmt.Errorf("failed to fetch schedules: %w", err),
		)
	}
	return schedules, nil
}

func (s *Store) AddToSchedule(
	ctx context.Context,
	channelID string,
	users []*entity.User,
) (AddResult, error) {

	// failure helpers
	var (
		failure  = failure[AddResult]
		rollback = rollback[AddResult]
	)

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return failure(err)
	}

	builders := asUserBuilders(tx, users...)
	users, err = tx.User.
		CreateBulk(builders...).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	schedule, err := tx.Schedule.Query().
		Where(schedule.SlackChannelIDEQ(channelID)).
		WithShifts(func(q *entity.ShiftQuery) {
			q.Order(entity.Asc(shift.FieldSequenceID))
		}).
		Only(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	var (
		active *entity.User
		shifts = asShifts(tx, schedule, users)
	)

	if setActiveShift(schedule) {
		active = users[0]
		shifts[0].SetStartedAt(time.Now())
		shifts[0].SetStatus(shift.StatusActive)
	}

	newShifts, err := tx.Shift.
		CreateBulk(shifts...).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return rollback(tx, err)
	}

	schedule.Edges.Shifts = append(schedule.Edges.Shifts, newShifts...)
	return AddResult{
		Schedule:       *schedule,
		NewActiveShift: active,
	}, nil
}
