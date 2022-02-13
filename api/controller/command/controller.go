package command

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/api/middleware"
	"github.com/syllabix/oncall/api/rest"
	"github.com/syllabix/oncall/slack/command"
	"go.uber.org/zap"
)

type Controller struct {
	verifier *middleware.SlackVerifier
	handler  command.Handler

	log *zap.Logger
}

func NewController(
	verifier *middleware.SlackVerifier,
	handler command.Handler,
	log *zap.Logger,
) rest.Controller {
	return rest.MakeController(
		&Controller{verifier, handler, log},
	)
}

// `/schedule` - manage an call schedules
// `/swap` - swap shifts between two members
// `/add` - adds a user to an on call rotation
// `/remove` - removes a user from the on call rotation
// `/override` - create an override for the current shift

func (ctrl *Controller) Register(router *httprouter.Router) {
	router.POST("/slack/command/:command", ctrl.HandleCommand)
}

func (ctrl *Controller) HandleCommand(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	cmd, err := slack.SlashCommandParse(r)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	res, err := ctrl.handler.Handle(cmd)
	if err != nil {
		ctrl.log.Error("command handler failed", zap.Error(err))
		switch err {
		case command.ErrInvalidSlackToken:
			http.Error(w,
				http.StatusText(http.StatusBadRequest),
				http.StatusBadRequest)

		case command.ErrUnsupportedCommand:
			http.Error(w,
				"Sorry... I don't know how to do that yet",
				http.StatusBadRequest)

		default:
			http.Error(w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		ctrl.log.Error("failed to encode and send back response to command",
			zap.Error(err),
			zap.String("command", cmd.Command),
		)
	}
}
