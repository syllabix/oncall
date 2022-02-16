package schedule

import (
	"github.com/google/uuid"
	"github.com/syllabix/oncall/datastore/model"
)

func shiftFromUser(scheduleID string, user model.User) model.Shift {
	return model.Shift{
		ID:         uuid.NewString(),
		UserID:     user.ID,
		ScheduleID: scheduleID,
	}
}
