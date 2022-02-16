package schedule

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syllabix/oncall/datastore/db"
	"github.com/syllabix/oncall/datastore/model"
)

type Store struct {
	db            db.Database
	create        *sqlx.NamedStmt
	listByChannel *sqlx.Stmt
}

func NewStore(db db.Database) (Store, error) {
	create, err := createStmt(db)
	if err != nil {
		return Store{}, fmt.Errorf("failed to setup schedule data store: %w", err)
	}

	listByChannel, err := db.Preparex(`
		SELECT * FROM schedules WHERE slack_channel_id = $1
	`)
	if err != nil {
		return Store{}, fmt.Errorf("failed to setup schedule data store: %w", err)
	}

	return Store{db, create, listByChannel}, nil
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

func (s *Store) GetForChannel(id string) (model.ScheduleSlice, error) {
	var schedules model.ScheduleSlice
	err := s.listByChannel.Select(&schedules, id)
	if err != nil {
		return failure[model.ScheduleSlice](
			fmt.Errorf("could not create schedule: %w", err),
		)
	}
	return schedules, nil
}

func (s *Store) AddToSchedule(channelID string, users []model.User) (Overview, error) {

	// failure helpers
	var (
		failure  = failure[Overview]
		rollback = rollback[Overview]
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

	var shifts model.ShiftSlice
	err = tx.Select(&shifts, "SELECT * FROM shifts WHERE schedule_id = $1", schedule.ID)
	if err != nil {
		return rollback(tx, err)
	}

	_, err = tx.NamedExec(upsertUser, users)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return failure(err)
	}

	return Overview{}, nil
}
