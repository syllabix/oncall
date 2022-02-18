package event

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var (
	ErrUnsupportedEvent = errors.New("the provided event is not supported")
)

type Handler interface {
	Handle(slackevents.EventsAPIEvent) error
}

type handler struct {
	client *slack.Client
}

func (h handler) Handle(event slackevents.EventsAPIEvent) error {

	switch ev := event.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		return h.handleAppMention(ev)

	default:
		return fmt.Errorf("%w: event type: %s", ErrUnsupportedEvent, event.Type)
	}

}

func (h handler) handleAppMention(event *slackevents.AppMentionEvent) error {

	// Grab the user name based on the ID of the one who mentioned the bot
	user, err := h.client.GetUserInfo(event.User)
	if err != nil {
		return err
	}

	// Check if the user said Hello to the bot
	text := strings.ToLower(event.Text)

	// Create the attachment and assigned based on the message
	attachment := slack.Attachment{}
	// Add Some default context like user who mentioned the bot
	attachment.Fields = []slack.AttachmentField{
		{
			Title: "Date",
			Value: time.Now().String(),
		}, {
			Title: "Initializer",
			Value: user.Name,
		},
	}
	if strings.Contains(text, "hello") {
		// Greet the user
		attachment.Text = fmt.Sprintf("Hello %s", user.Name)
		attachment.Pretext = "Greetings"
		attachment.Color = "#4af030"
	} else {
		// Send a message to the user
		attachment.Text = fmt.Sprintf("How can I help you %s?", user.Name)
		attachment.Pretext = "How can I be of service"
		attachment.Color = "#3d3d3d"
	}

	// Send the message to the channel
	// The Channel is available in the event message
	_, _, err = h.client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}
