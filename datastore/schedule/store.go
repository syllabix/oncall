package schedule

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syllabix/oncall/datastore/db"
	"github.com/syllabix/oncall/datastore/model"
)

type Store struct {
	create        *sqlx.NamedStmt
	listByChannel *sqlx.Stmt
}

func NewStore(db db.Database) (Store, error) {
	create, err := db.PrepareNamed(`
		INSERT INTO oncall_schedule 
		(team_slack_id, name, interval, start_time, end_time, slack_channel_id) VALUES
		(:team_slack_id, :name, :interval, :start_time, :end_time, :slack_channel_id)
		returning id, created_at, updated_at`,
	)
	if err != nil {
		return Store{}, fmt.Errorf("failed to setup schedule data store: %w", err)
	}

	listByChannel, err := db.Preparex(`
		SELECT * FROM oncall_schedule WHERE slack_channel_id = $1
	`)
	if err != nil {
		return Store{}, fmt.Errorf("failed to setup schedule data store: %w", err)
	}

	return Store{create, listByChannel}, nil
}

func (s *Store) Create(schedule model.OncallSchedule) (model.OncallSchedule, error) {
	err := s.create.Get(&schedule, &schedule)
	if err != nil {
		return model.OncallSchedule{}, fmt.Errorf("failed to create new schedule: %w", err)
	}
	return schedule, nil
}

func (s *Store) GetForChannel(id string) (model.OncallScheduleSlice, error) {
	var schedules model.OncallScheduleSlice
	err := s.listByChannel.Select(&schedules, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedules for channel id: %w", err)
	}
	return schedules, nil
}
