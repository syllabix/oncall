package schedule

import (
	"github.com/syllabix/oncall/datastore/model"
)

type AddResult struct {
	Schedule       model.Schedule
	NewActiveShift *model.User
}
