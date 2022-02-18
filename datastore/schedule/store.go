package schedule

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syllabix/oncall/datastore/db"
	"github.com/syllabix/oncall/datastore/model"
	"github.com/volatiletech/null/v8"
)

type Store struct {
	db        db.Database
	create    *sqlx.NamedStmt
	update    *sqlx.NamedStmt
	schedules *sqlx.Stmt
}

func NewStore(db db.Database) (Store, error) {
	create, err := createStmt(db)
	if err != nil {
		return Store{}, fmt.Errorf("failed to setup schedule data store: %w", err)
	}

	schedules, err := db.Preparex(`
		SELECT * FROM schedules WHERE is_enabled = true
	`)
	if err != nil {
		return Store{}, fmt.Errorf("failed to setup schedule data store: %w", err)
	}

	update, err := updateStmt(db)
	if err != nil {
		return Store{}, fmt.Errorf("failed to setup schedule data store: %w", err)
	}

	return Store{db, create, update, schedules}, nil
}

func (s *Store) Create(schedule model.Schedule) (model.Schedule, error) {
	err := s.create.Get(&schedule, &schedule)
	if err != nil {
		return failure[model.Schedule](
			fmt.Errorf("could not create schedule: %w", err),
		)
	}

	return schedule, nil
}

func (s *Store) Update(schedule model.Schedule) (model.Schedule, error) {
	err := s.update.Get(&schedule, &schedule)
	if err != nil {
		return failure[model.Schedule](
			fmt.Errorf("could not update schedule: %w", err),
		)
	}
	return schedule, nil
}

func (s *Store) GetByID(scheduleID string) (model.Schedule, error) {
	var sched model.Schedule
	err := s.db.Get(&sched, "SELECT * FROM schedules WHERE id = $1", scheduleID)
	if err != nil {
		return sched, err
	}
	return sched, err
}

func (s *Store) GetEnabledSchedules() (model.ScheduleSlice, error) {
	var schedules model.ScheduleSlice
	err := s.schedules.Select(&schedules)
	if err != nil {
		return failure[model.ScheduleSlice](
			fmt.Errorf("failed to fetch schedules: %w", err),
		)
	}
	return schedules, nil
}

// TODO: this is pretty sad - but should work for now - rethink the data model and reduce
// db trips if possible
func (s *Store) AddToSchedule(channelID string, users []model.User) (AddResult, error) {

	// failure helpers
	var (
		failure  = failure[AddResult]
		rollback = rollback[AddResult]
	)

	tx, err := s.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return failure(err)
	}

	var schedule model.Schedule
	err = tx.Get(&schedule, "SELECT * FROM schedules WHERE slack_channel_id = $1", channelID)
	if err != nil {
		return rollback(tx, err)
	}

	_, err = tx.NamedExec(upsertUser, users)
	if err != nil {
		return rollback(tx, err)
	}

	var (
		active *model.User
	)

	if schedule.ActiveShift.IsZero() {
		active = &users[0]
		schedule.ActiveShift = null.StringFrom(active.ID)
	}

	schedule.Shifts = append(schedule.Shifts, uniqueIdsFrom(users, schedule)...)
	_, err = tx.NamedExec(updateScheduleShift, schedule)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return failure(err)
	}

	return AddResult{
		Schedule:       schedule,
		NewActiveShift: active,
	}, nil
}
