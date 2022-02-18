package event

import (
	"errors"
	"fmt"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/syllabix/oncall/service/schedule"
)

var (
	ErrUnsupportedEvent = errors.New("the provided event is not supported")
)

type Handler interface {
	Handle(slackevents.EventsAPIEvent) error
}

type handler struct {
	client  *slack.Client
	manager schedule.Manager
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

	shift, err := h.manager.GetActiveShift(event.Channel)
	if err != nil {
		_, _, err = h.client.PostMessage(
			event.Channel,
			slack.MsgOptionText(
				":fearful: no one is on call at the moment...",
				false,
			),
			slack.MsgOptionTS(event.TimeStamp),
		)
		return nil
	}

	_, _, err = h.client.PostMessage(
		event.Channel,
		slack.MsgOptionText(
			fmt.Sprintf("paging %s", shift.SlackHandle),
			false,
		),
		slack.MsgOptionTS(event.TimeStamp),
	)
	if err != nil {
		return fmt.Errorf("something went wrong posting the on call reply: %v", err)
	}
	return nil
}
