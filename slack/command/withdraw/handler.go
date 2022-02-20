package withdraw

import "github.com/slack-go/slack"

type Handler interface {
	Handle(slack.SlashCommand) (slack.Attachment, error)
}

func NewHandler() Handler

type handler struct {
}

func (h *handler) Handle(slack.SlashCommand) (slack.Attachment, error) {
	return slack.Attachment{}, nil
}
