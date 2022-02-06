package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/api/middleware"
	"github.com/syllabix/oncall/api/rest"
	"github.com/syllabix/oncall/service/schedule"
	"go.uber.org/zap"
)

type Controller struct {
	verifier  *middleware.SlackVerifier
	scheduler schedule.Manager

	log *zap.Logger
}

func NewController(
	verifier *middleware.SlackVerifier,
	scheduler schedule.Manager,
	log *zap.Logger,
) rest.Controller {
	return rest.MakeController(
		&Controller{verifier, scheduler, log},
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
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadGateway)
		return
	}

	cmd.ValidateToken()

	w.Header().Set("Content-Type", "application/json")

	switch cmd.Command {
	case "/schedule":
		res := ctrl.scheduler.Prepare()
		output, _ := json.Marshal(&res)
		fmt.Println(string(output))
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			ctrl.log.Error("failed to to had")
		}
		return

	default:
		http.Error(w, "Sorry... I don't know how to do that yet", http.StatusBadGateway)
	}
}
