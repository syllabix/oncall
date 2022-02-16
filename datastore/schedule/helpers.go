package schedule

import "github.com/syllabix/oncall/datastore/model"

func uniqueIdsFrom(users []model.User, in model.Schedule) (ids []string) {

	set := make(map[string]struct{})

	for _, id := range in.Shifts {
		set[id] = struct{}{}
	}

	for _, user := range users {
		_, contains := set[user.ID]
		if contains {
			continue
		}

		ids = append(ids, user.ID)
	}
	return
}
