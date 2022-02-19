package add

import (
	"context"
	"errors"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/datastore/model"
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/slack/util/mention"
)

var (
	ErrInvalidInput        = errors.New("invalid input for add - one or more @ user mentions are required")
	ErrNoAvailableSchedule = errors.New("there is no available schedule to add a user to")
	ErrUserAlreadyAdded    = errors.New("one or more of the provided users are already added to the schedule")
)

type Result struct {
	ScheduleName   string
	EngineerOnDuty string
}

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

	var usrModels model.UserSlice
	for _, user := range users {
		usr, err := h.client.GetUserInfo(user.ID)
		if err != nil {
			return err
		}
		usrModels = append(usrModels, toUserModel(usr, user.HandleID))
	}

	result, err := h.scheduler.AddToSchedule(context.TODO(), cmd.ChannelID, usrModels)
	switch {
	case err == nil:
		if result.NewActiveShift != nil {
			// ignore errors for now - perhaps retry if this fails?
			h.client.SetTopicOfConversation(cmd.ChannelID,
				result.Schedule.Name+": "+result.NewActiveShift.SlackHandle,
			)
		}
		return nil

	case errors.Is(err, schedule.ErrNotFound):
		return ErrNoAvailableSchedule

	case errors.Is(err, schedule.ErrAlreadyInUse):
		return ErrUserAlreadyAdded

	default:
		return err
	}
}
