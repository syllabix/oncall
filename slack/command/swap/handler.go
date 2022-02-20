package swap

import (
	"context"
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
	if err != nil || len(mentions) != 2 {
		return slack.Attachment{
			Title: "Something went wrong",
			Text:  ":warning: please @ mention two team members that would like to swap shifts",
		}, nil
	}

	nextActive, err := h.manager.SwapShift(context.TODO(),
		shift.SwapRequest{
			ChannelID:    cmd.ChannelID,
			UserSlackID1: mentions[0].ID,
			UserSlackID2: mentions[1].ID,
		})
	if err != nil {
		h.log.Error("failed to complete swap", zap.Error(err))
		return slack.Attachment{
			Title: "Something went wrong",
			Text:  ":warning: Wow... so sorry but I was unable to complete the swap. Please make sure both team members are assigned and try again.",
		}, nil
	}

	if nextActive != nil {
		h.client.SetTopicOfConversation(cmd.ChannelID,
			fmt.Sprintf("%s: %s", nextActive.ScheduleName, nextActive.SlackHandle),
		)
	}

	return slack.Attachment{
		Title: "Team member withdrawn ok",
		Text:  ":white_check_mark: team members have had there shifts swapped successfully",
	}, nil

}
