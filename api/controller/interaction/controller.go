package interaction

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/api/middleware"
	"github.com/syllabix/oncall/api/rest"
	"github.com/syllabix/oncall/service/schedule"
	"github.com/syllabix/oncall/slack/interaction"
	"go.uber.org/zap"
)

type Controller struct {
	handler   interaction.Handler
	verifier  *middleware.SlackVerifier
	scheduler schedule.Manager

	log *zap.Logger
}

func NewController(
	handler interaction.Handler,
	verifier *middleware.SlackVerifier,
	scheduler schedule.Manager,
	log *zap.Logger,
) rest.Controller {
	return rest.MakeController(
		&Controller{handler, verifier, scheduler, log},
	)
}

func (ctrl *Controller) Register(router *httprouter.Router) {
	router.POST("/slack/interaction", ctrl.HandleInteraction)
}

func (ctrl *Controller) HandleInteraction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	payload := r.FormValue("payload")
	if len(payload) < 1 {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	var callback slack.InteractionCallback
	err = json.Unmarshal([]byte(payload), &callback)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	err = ctrl.handler.Handle(callback)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	// ctrl.scheduler.Create(schedule)
	w.WriteHeader(200)
	fmt.Fprint(w, "Ok")
}
