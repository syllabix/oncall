package slack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/syllabix/oncall/api/middleware"
	"github.com/syllabix/oncall/api/rest"
	"github.com/syllabix/oncall/service/slack/event"
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
	router.POST("/slack/interaction", ctrl.HandleInteraction)
	router.POST("/slack/load-options", ctrl.verifier.Verify(ctrl.LoadOptions))
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
			ctrl.log.Error("failed to unmarshal url verification body", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprint(w, "Ok")
}

func (ctrl *Controller) HandleInteraction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	var callback slack.InteractionCallback
	err = json.Unmarshal([]byte(r.Form.Get("payload")), &callback)

	fmt.Printf("%+v", callback)

	w.WriteHeader(200)
	fmt.Fprint(w, "Ok")
}

func (ctrl *Controller) LoadOptions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var payload strings.Builder
	_, err := io.Copy(&payload, r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println("load options", payload.String())
	w.WriteHeader(200)
	fmt.Fprint(w, "Ok")
}
