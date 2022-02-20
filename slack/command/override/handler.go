package override

import (
	"context"
	"errors"
	"fmt"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/service/shift"
	"github.com/syllabix/oncall/slack/util/mention"
	"go.uber.org/zap"
)

type Handler interface {
	Handle(slack.SlashCommand) (slack.Attachment, error)
}

func NewHandler(
	client *slack.Client,
	manager shift.Manager,
	parser *mention.Parser,
	log *zap.Logger,
) Handler {
	return &handler{client, manager, parser, log}
}

type handler struct {
	client  *slack.Client
	manager shift.Manager
	parser  *mention.Parser

	log *zap.Logger
}

func (h *handler) Handle(cmd slack.SlashCommand) (slack.Attachment, error) {
	mentions, err := h.parser.ParseList(cmd.Text)
	if err != nil || len(mentions) != 1 {
		return slack.Attachment{
			Title: "Something went wrong",
			Text:  ":warning: please @ mention one team member that will override the current shift",
		}, nil
	}

	overrideShift, err := h.manager.ApplyOverride(context.TODO(),
		shift.OverrideRequest{
			ChannelID:   cmd.ChannelID,
			UserSlackID: mentions[0].ID,
		})
	if err != nil {
		if errors.Is(err, shift.ErrNotFound) {
			return slack.Attachment{
				Title: "Shift not found",
				Text:  ":warning: The team member is not currently assigned to this schedule",
			}, nil
		}

		h.log.Error("failed to remove shift", zap.Error(err))
		return slack.Attachment{
			Title: "Something went wrong",
			Text:  ":cry: O wow this is embarrassing... but I was not able to set up the override. Try again?",
		}, nil
	}

	h.client.SetTopicOfConversation(cmd.ChannelID,
		fmt.Sprintf("%s: %s (override)", overrideShift.ScheduleName, overrideShift.SlackHandle),
	)

	return slack.Attachment{
		Title: "Schedule override applied",
		Text:  ":white_check_mark: The override has been set up ok",
	}, nil

}
