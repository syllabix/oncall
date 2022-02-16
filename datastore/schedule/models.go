package schedule

import (
	"github.com/syllabix/oncall/datastore/model"
)

type Overview struct {
	Schedule    model.Schedule
	ActiveShift model.Shift
}
