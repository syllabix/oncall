package withdraw

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
	users, err := h.parser.ParseList(cmd.Text)
	if err != nil {
		return slack.Attachment{
			Title: "Something went wrong",
			Text:  ":warning: please @ mention a single team member to withdraw them from the schedule",
		}, nil
	}

	if len(users) > 1 {
		return slack.Attachment{
			Title: "Something went wrong",
			Text:  ":warning: I can only withdraw one team member from the schedule at a time",
		}, nil
	}

	nextActive, err := h.manager.RemoveShift(context.TODO(),
		shift.RemoveRequest{
			ChannelID:   cmd.ChannelID,
			UserSlackID: users[0].ID,
		})
	if err != nil {
		if errors.Is(err, shift.ErrNotFound) {
			return slack.Attachment{
				Title: "Shift not found",
				Text:  ":warning: The team member is not currently assigned to this schedule",
			}, nil
		}

		var emptyErr *shift.ErrScheduleEmpty
		if errors.As(err, &emptyErr) {
			h.client.SetTopicOfConversation(cmd.ChannelID,
				fmt.Sprintf("%s: :sleeping:", emptyErr.Name),
			)

			return slack.Attachment{
				Title: "Shift not found",
				Text:  ":warning: There are no longer any shifts assigned to this schedule",
			}, nil
		}

		h.log.Error("failed to remove shift", zap.Error(err))
		return slack.Attachment{
			Title: "Something went wrong",
			Text:  ":cry: O wow this is embarrassing... but I was not able to withdraw the team member from the schedule. Try again?",
		}, nil
	}

	if nextActive != nil {
		h.client.SetTopicOfConversation(cmd.ChannelID,
			fmt.Sprintf("%s: %s", nextActive.ScheduleName, nextActive.SlackHandle),
		)
	}

	return slack.Attachment{
		Title: "Team member withdrawn ok",
		Text:  ":white_check_mark: the team member has been withdrawn from the schedule",
	}, nil
}
