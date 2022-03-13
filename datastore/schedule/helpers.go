package schedule

import (
	"github.com/syllabix/oncall/common/is"
	"github.com/syllabix/oncall/datastore/entity"
	"github.com/syllabix/oncall/datastore/entity/shift"
)

func setActiveShift(schedule *entity.Schedule) bool {
	if schedule.WeekdaysOnly && is.TheWeekend() {
		return false
	}

	for _, s := range schedule.Edges.Shifts {
		if s.Status == shift.StatusActive {
			return false
		}
	}

	return true
}

func asShifts(tx *entity.Tx, schedule *entity.Schedule, users []*entity.User) (shifts []*entity.ShiftCreate) {
	for i := range users {
		user := users[i]
		shifts = append(shifts,
			tx.Shift.Create().
				SetUser(user).
				SetSchedule(schedule),
		)
	}
	return shifts
}

func asUserBuilders(tx *entity.Tx, users ...*entity.User) (builders []*entity.UserCreate) {
	for i := range users {
		user := users[i]
		builders = append(builders,
			tx.User.Create().
				SetSlackID(user.SlackID).
				SetSlackHandle(user.SlackHandle).
				SetEmail(user.Email).
				SetFirstName(user.FirstName).
				SetLastName(user.LastName).
				SetAvatarURL(user.AvatarURL).
				SetDisplayName(user.DisplayName),
		)
	}

	return
}
