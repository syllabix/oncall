package add

import (
	"errors"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/slack/util/mention"
)

var (
	ErrInvalidInput        = errors.New("invalid input for add - one or more @ user mentions are required")
	ErrNoAvailableSchedule = errors.New("there is no available schedule to add a user to")
)

type Handler interface {
	AddToSchedule(slack.SlashCommand) error
}

func NewHandler(client *slack.Client, parser *mention.Parser, scheduler schedule.Store) Handler {
	return &handler{parser, client, scheduler}
}

type handler struct {
	parser    *mention.Parser
	client    *slack.Client
	scheduler schedule.Store
}

func (h *handler) AddToSchedule(cmd slack.SlashCommand) error {
	users, err := h.parser.ParseList(cmd.Text)
	if err != nil {
		return ErrInvalidInput
	}

	var usrModels []model.User
	for _, user := range users {
		usr, err := h.client.GetUserInfo(user.ID)
		if err != nil {
			return err
		}
		usrModels = append(usrModels, toUserModel(usr))
	}

	_, err = h.scheduler.AddToSchedule(cmd.ChannelID, usrModels)
	if err != nil {
		if errors.Is(err, schedule.ErrNotFound) {
			return ErrNoAvailableSchedule
		}
		return err
	}

	return nil
}
