package event

import (
	"github.com/slack-go/slack/slackevents"
	"go.uber.org/zap"
)

// handlermonitor is an observability wrapper around
// an event Handler. It's primary purpose is to decouple
// monitoring code from actual business logic
type handlermonitor struct {
	handler Handler
	log     *zap.Logger
}

func (h *handlermonitor) Handle(event slackevents.EventsAPIEvent) error {
	err := h.handler.Handle(event)
	if err != nil {
		h.log.Error("event handling failed", zap.Error(err))
	}
	return err
}
