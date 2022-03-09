package event

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/syllabix/oncall/service/schedule"
	"go.uber.org/zap"
)

var (
	ErrUnsupportedEvent = errors.New("the provided event is not supported")
)

type Handler interface {
	Handle(slackevents.EventsAPIEvent) error
}

func NewHandler(client *slack.Client, manager schedule.Manager, log *zap.Logger) Handler {
	return &handlermonitor{
		handler: &handler{client, manager},
		log:     log,
	}
}

type handler struct {
	client  *slack.Client
	manager schedule.Manager
}

func (h handler) Handle(event slackevents.EventsAPIEvent) error {

	switch ev := event.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		return h.handle(ev)

	default:
		return fmt.Errorf("%w: event type: %s", ErrUnsupportedEvent, event.Type)
	}

}

func (h handler) handle(event *slackevents.AppMentionEvent) error {

	var (
		response string
	)

	shift, err := h.manager.GetActiveShift(event.Channel)
	switch {
	case err == nil:
		response = "paging " + shift.SlackHandle

	case errors.Is(err, schedule.ErrNoActiveShift):
		response = ":sleeping: no one is on call at the moment"

	default:
		response = ":cold_sweat: well this is embarrassing, but I am having a little trouble at the moment."
	}

	_, _, postErr := h.client.PostMessage(
		event.Channel,
		slack.MsgOptionText(response, false),
		slack.MsgOptionTS(event.TimeStamp),
	)
	if postErr != nil {
		return errors.WithStack(postErr)
	}

	return err
}
