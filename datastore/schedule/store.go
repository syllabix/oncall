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
		SELECT * FROM oncall_schedule WHERE slack_channel_id = $1
	`)
	if err != nil {
		return Store{}, fmt.Errorf("failed to setup schedule data store: %w", err)
	}

	return Store{db, create, listByChannel}, nil
}

func (s *Store) Create(schedule model.OncallSchedule) (model.OncallSchedule, error) {
	err := s.create.Get(&schedule, &schedule)
	if err != nil {
		return failure[model.OncallSchedule](
			fmt.Errorf("could not create schedule: %w", err),
		)
	}
	return schedule, nil
}

func (s *Store) GetForChannel(id string) (model.OncallScheduleSlice, error) {
	var schedules model.OncallScheduleSlice
	err := s.listByChannel.Select(&schedules, id)
	if err != nil {
		return failure[model.OncallScheduleSlice](
			fmt.Errorf("could not create schedule: %w", err),
		)
	}
	return schedules, nil
}

func (s *Store) AddToSchedule(channelID string, users model.UserSlice) (model.OncallSchedule, error) {

	// failure helpers
	var (
		failure  = failure[model.OncallSchedule]
		rollback = rollback[model.OncallSchedule]
	)

	tx, err := s.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return failure(err)
	}

	var schedule model.OncallSchedule
	err = tx.Get(&schedule, "SELECT * WHERE slack_channel_id = $1", channelID)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return failure(err)
	}
}
