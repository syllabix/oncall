package add

import (
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/datastore/entity"
)

func toUserModel(user *slack.User, handle string) *entity.User {
	return &entity.User{
		SlackID:     user.ID,
		SlackHandle: handle,
		Email:       user.Profile.Email,
		FirstName:   user.Profile.FirstName,
		LastName:    user.Profile.LastName,
		DisplayName: user.Profile.DisplayNameNormalized,
		AvatarURL:   user.Profile.Image192,
	}
}
