package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/syllabix/oncall/common/db"
	"github.com/syllabix/oncall/datastore/model"
)

type Store struct {
	db db.Postgres
}

func NewStore(db db.Postgres) (Store, error) {
	return Store{db}, nil
}

func (s *Store) Create(ctx context.Context, schedule *model.Schedule) (*model.Schedule, error) {
	err := schedule.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return failure[*model.Schedule](
			fmt.Errorf("could not create schedule: %w", err),
		)
	}

	return schedule, nil
}

func (s *Store) Update(ctx context.Context, schedule *model.Schedule) (*model.Schedule, error) {
	_, err := schedule.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return failure[*model.Schedule](
			fmt.Errorf("could not update schedule: %w", err),
		)
	}
	return schedule, nil
}

func (s *Store) GetByID(ctx context.Context, id int) (*model.Schedule, error) {
	sched, err := model.Schedules(
		model.ScheduleWhere.ID.EQ(id),
		qm.Load(
			model.ScheduleRels.Shifts,
			qm.OrderBy(model.ShiftColumns.SequenceID),
		),
		qm.Load(qm.Rels(
			model.ScheduleRels.Shifts,
			model.ShiftRels.User,
		)),
	).One(ctx, s.db)
	if err != nil {
		return failure[*model.Schedule](
			fmt.Errorf("could not get schedule by id: %w", err),
		)
	}
	return sched, nil
}

func (s *Store) GetByChannelID(ctx context.Context, id string) (*model.Schedule, error) {
	sched, err := model.Schedules(
		model.ScheduleWhere.SlackChannelID.EQ(id),
		qm.Load(
			model.ScheduleRels.Shifts,
			qm.OrderBy(model.ShiftColumns.SequenceID),
		),
		qm.Load(qm.Rels(
			model.ScheduleRels.Shifts,
			model.ShiftRels.User,
		)),
	).One(ctx, s.db)
	if err != nil {
		return failure[*model.Schedule](
			fmt.Errorf("could not get schedule by channel id: %w", err),
		)
	}
	return sched, nil
}

func (s *Store) GetEnabledSchedules(ctx context.Context) (model.ScheduleSlice, error) {
	schedules, err := model.Schedules(
		model.ScheduleWhere.IsEnabled.EQ(true),
		qm.Load(
			model.ScheduleRels.Shifts,
			qm.OrderBy(model.ShiftColumns.SequenceID),
		),
		qm.Load(qm.Rels(
			model.ScheduleRels.Shifts,
			model.ShiftRels.User,
		)),
	).All(ctx, s.db)
	if err != nil {
		return failure[model.ScheduleSlice](
			fmt.Errorf("failed to fetch schedules: %w", err),
		)
	}
	return schedules, nil
}

func (s *Store) AddToSchedule(
	ctx context.Context,
	channelID string,
	users model.UserSlice,
) (AddResult, error) {

	// failure helpers
	var (
		failure  = failure[AddResult]
		rollback = rollback[AddResult]
	)

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return failure(err)
	}

	schedule, err := model.Schedules(
		model.ScheduleWhere.SlackChannelID.EQ(channelID),
		qm.Load(model.ScheduleRels.Shifts),
	).One(ctx, tx)
	if err != nil {
		return rollback(tx, err)
	}

	rows, err := tx.NamedQuery(upsertStmt, users)
	if err != nil {
		return rollback(tx, err)
	}

	i := 0
	for rows.Next() {
		user := users[i]
		err = rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return rollback(tx, err)
		}
		i++
	}

	err = rows.Close()
	if err != nil {
		return rollback(tx, err)
	}

	var (
		active *model.User
		shifts = asShifts(schedule.ID, users)
	)

	if setActiveShift(schedule) {
		active = users[0]
		shifts[0].StartedAt = null.TimeFrom(time.Now())
		shifts[0].Status = null.StringFrom(model.ShiftStatusActive)
	}

	err = schedule.AddShifts(ctx, tx, true, shifts...)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return rollback(tx, err)
	}

	return AddResult{
		Schedule:       *schedule,
		NewActiveShift: active,
	}, nil
}
