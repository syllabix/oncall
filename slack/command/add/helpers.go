package add

import (
	"github.com/google/uuid"
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/datastore/model"
)

func toUserModel(user *slack.User) model.User {
	return model.User{
		ID:          uuid.NewString(),
		SlackID:     user.ID,
		Email:       user.Profile.Email,
		FirstName:   user.Profile.FirstName,
		LastName:    user.Profile.LastName,
		DisplayName: user.Profile.DisplayNameNormalized,
		AvatarURL:   user.Profile.Image192,
	}
}
