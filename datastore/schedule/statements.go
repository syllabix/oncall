package schedule

import (
	"github.com/jmoiron/sqlx"
	"github.com/syllabix/oncall/datastore/db"
)

func createStmt(db db.Database) (*sqlx.NamedStmt, error) {
	return db.PrepareNamed(`
	INSERT INTO schedules 
	(team_slack_id, name, interval, start_time, end_time, slack_channel_id) VALUES
	(:team_slack_id, :name, :interval, :start_time, :end_time, :slack_channel_id)
	RETURNING id, created_at, updated_at`,
	)
}

const upsertUser = `
	INSERT INTO users (id, slack_id, slack_handle, email, first_name, last_name, avatar_url, display_name)
	VALUES (:id, :slack_id, :slack_handle, :email, :first_name, :last_name, :avatar_url, :display_name)
	ON CONFLICT (slack_id) DO NOTHING`

const updateScheduleShift = `
	UPDATE schedules set shifts = :shifts, active_shift = :active_shift, updated_at = now() where id = :id;
`
