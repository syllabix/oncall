package slack

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/slack-go/slack/slackevents"
	"github.com/syllabix/oncall/api/middleware"
	"github.com/syllabix/oncall/api/rest"
	"github.com/syllabix/oncall/slack/event"
	"go.uber.org/zap"
)

type Controller struct {
	parser   event.Parser
	handler  event.Handler
	verifier *middleware.SlackVerifier

	log *zap.Logger
}

func NewController(
	handler event.Handler,
	parser event.Parser,
	verifier *middleware.SlackVerifier,
	log *zap.Logger,
) rest.Controller {
	return rest.MakeController(
		&Controller{parser, handler, verifier, log},
	)
}

func (ctrl *Controller) Register(router *httprouter.Router) {
	router.POST("/slack/action", ctrl.HandleAction)
}

func (ctrl *Controller) HandleAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	event, err := ctrl.parser.Parse(r.Body)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	switch event.Type {
	case slackevents.URLVerification:
		ctrl.verifyURL(w, event.RawMessage)
		return

	case slackevents.CallbackEvent:
		err = ctrl.handler.Handle(event.EventsAPIEvent)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprint(w, "Ok")
}
