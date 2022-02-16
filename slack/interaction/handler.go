package interaction

import (
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/slack/interaction/schedule"
)

// A Handler is responsible for dealing with slack interaction callbacks
type Handler interface {

	// Handle the provided callback message from slack, returning an
	// error in the event there was one
	Handle(slack.InteractionCallback) error
}

func NewHandler(
	scheduler schedule.Creator,
) Handler {
	return &handler{scheduler}
}

type handler struct {
	scheduler schedule.Creator
}

func (h *handler) Handle(callback slack.InteractionCallback) error {

	switch actionTypeFor(callback) {
	case "create_new_schedule_action_submit":
		return h.scheduler.Create(callback)

	// otherwise - it is an action type we do not handle
	// this is effectively a no-op and should not error out
	default:
		return nil
	}

}
