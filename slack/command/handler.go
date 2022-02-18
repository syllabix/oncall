package command

import (
	"errors"
	"strings"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/config"
	"github.com/syllabix/oncall/slack/command/add"
	"github.com/syllabix/oncall/slack/command/schedule"
	"go.uber.org/zap"
)

var (
	ErrInvalidSlackToken  = errors.New("the token associated with command is not valid and cannot be trusted")
	ErrUnsupportedCommand = errors.New("the provided command is not supported")
)

// Handler is responsible for dealing with slash commands
// sent by slack. The result of a successful handle returns
// a data that can be returned to Slack
type Handler interface {
	Handle(slack.SlashCommand) (any, error)
}

func NewHandler(
	config config.SlackSettings,
	adder add.Handler,
	schedule schedule.Handler,
	log *zap.Logger,
) Handler {
	return &handler{config, adder, schedule, log}
}

type handler struct {
	config   config.SlackSettings
	adder    add.Handler
	schedule schedule.Handler

	log *zap.Logger
}

func (h *handler) Handle(cmd slack.SlashCommand) (any, error) {

	valid := cmd.ValidateToken(h.config.VerificationToken)
	if !valid {
		return nil, ErrInvalidSlackToken
	}

	switch cmd.Command {
	case "/schedule":

		subcommand := strings.TrimSpace(strings.ToLower(cmd.Text))

		switch subcommand {
		case schedule.Create:
			return schedule.NewForm(), nil

		case schedule.View:
			return h.schedule.GetUpcomingShifts(cmd.ChannelID)

		// case schedule.Edit:
		default:
			return slack.Attachment{
				Title: "Unsupported slash sub command",
				Text:  ":warning: The */schedule* command expects either a *create*, *view*, or *edit* sub command",
			}, nil
		}

	case "/add":
		err := h.adder.AddToSchedule(cmd)
		if err != nil {
			h.log.Error("failed to add user to schedule",
				zap.Error(err),
				zap.String("channel", cmd.ChannelID))

			switch {
			case errors.Is(err, add.ErrInvalidInput):
				return slack.Attachment{
					Title: "Failed to create new shifts",
					Text:  ":warning: The */add* command expects one more @ user mentions to add them to a schedule",
				}, nil

			case errors.Is(err, add.ErrNoAvailableSchedule):
				return slack.Attachment{
					Title: "Failed to create new shifts",
					Text:  ":warning: There is no on call schedule set up for this channel. Please create one with the */schedule create* command",
				}, nil

			default:
				return slack.Attachment{
					Title: "Something went wrong",
					Text:  ":cry: O wow this is embarrassing but I was not able to add team members to the schedule",
				}, nil
			}

		}
		return slack.Attachment{
			Title: "Schedule updated!",
			Text:  ":white_check_mark: Your on call schedule has been updated",
		}, nil

	default:
		return slack.Attachment{
			Title: "Unsupported slash command",
			Text:  ":cry: Sorry - but I don't know how to handle that command yet...",
		}, nil

	}
}
