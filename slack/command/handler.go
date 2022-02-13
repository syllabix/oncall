package command

import (
	"errors"
	"strings"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/config"
	"github.com/syllabix/oncall/slack/command/schedule"
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
) Handler {
	return &handler{config}
}

type handler struct {
	config config.SlackSettings
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

		// case schedule.View:
		// case schedule.Edit:
		default:
			return slack.Attachment{
				Title: "Unsupported slash sub command",
				Text:  ":warning: The */schedule* command expects either a *create*, *view*, or *edit* sub command",
			}, nil
		}

	default:
		return slack.Attachment{
			Title: "Unsupported slash command",
			Text:  ":cry: Sorry - but I don't know how to handle that command yet...",
		}, nil

	}
}
