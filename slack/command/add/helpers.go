package add

import (
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/datastore/model"
)

func toUserModel(user *slack.User, handle string) *model.User {
	return &model.User{
		SlackID:     user.ID,
		SlackHandle: handle,
		Email:       user.Profile.Email,
		FirstName:   user.Profile.FirstName,
		LastName:    user.Profile.LastName,
		DisplayName: user.Profile.DisplayNameNormalized,
		AvatarURL:   user.Profile.Image192,
	}
}
