package add

import (
	"errors"
	"fmt"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/slack/util/mention"
)

var (
	ErrInvalidInput = errors.New("invalid input for add - one or more @ user mentions are required")
)

type Handler interface {
	AddToSchedule(slack.SlashCommand) error
}

func NewHandler(parser *mention.Parser, scheduler schedule.Store) Handler {
	return &handler{parser, scheduler}
}

type handler struct {
	parser    *mention.Parser
	scheduler schedule.Store
}

func (h *handler) AddToSchedule(cmd slack.SlashCommand) error {
	users, err := h.parser.ParseList(cmd.Text)
	if err != nil {
		return ErrInvalidInput
	}

	fmt.Println(users)
	return nil
}
