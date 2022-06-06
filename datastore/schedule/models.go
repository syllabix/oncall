package schedule

import (
	"github.com/syllabix/oncall/datastore/entity"
)

type AddResult struct {
	Schedule       entity.Schedule
	NewActiveShift *entity.User
}
