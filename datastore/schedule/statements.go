package schedule

import (
	"github.com/jmoiron/sqlx"
	"github.com/syllabix/oncall/datastore/db"
)

func createStmt(db db.Database) (*sqlx.NamedStmt, error) {
	return db.PrepareNamed(`
	INSERT INTO oncall_schedule 
	(team_slack_id, name, interval, start_time, end_time, slack_channel_id) VALUES
	(:team_slack_id, :name, :interval, :start_time, :end_time, :slack_channel_id)
	returning id, created_at, updated_at`,
	)
}
