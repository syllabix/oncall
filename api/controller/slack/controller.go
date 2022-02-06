package slack

import (
	"bytes"
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
	"go.uber.org/zap"
)

type Controller struct {
	client   *slack.Client
	verifier *middleware.SlackVerifier
	log      *zap.Logger
}

func NewController(
	client *slack.Client,
	verifier *middleware.SlackVerifier,
	log *zap.Logger,
) rest.Controller {
	return rest.MakeController(&Controller{client, verifier, log})
}

func (ctrl *Controller) Register(router *httprouter.Router) {
	router.POST("/slack/action", ctrl.verifier.Verify(ctrl.HandleAction))
	router.POST("/slack/interaction", ctrl.verifier.Verify(ctrl.HandleInteraction))
	router.POST("/slack/load-options", ctrl.verifier.Verify(ctrl.LoadOptions))
}

func (ctrl *Controller) HandleAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var payload bytes.Buffer
	_, err := io.Copy(&payload, r.Body)
	if err != nil {
		ctrl.log.Error("failed to read request body", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	event, err := slackevents.ParseEvent(payload.Bytes(), slackevents.OptionNoVerifyToken())
	if err != nil {
		ctrl.log.Error("failed to parse slack event", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	switch event.Type {
	case slackevents.URLVerification:
		var res slackevents.ChallengeResponse
		err := json.Unmarshal(payload.Bytes(), &res)
		if err != nil {
			ctrl.log.Error("failed to unmarshal url verification body", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			ctrl.log.Error("failed to write and encode url challenge", zap.Error(err))
			return
		}

	default:
		if err != nil {
			w.WriteHeader(500)
			return
		}
		fmt.Println("handle event", payload.String())

		w.WriteHeader(200)
		fmt.Fprint(w, "Ok")
	}

}

func (ctrl *Controller) HandleInteraction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var payload strings.Builder
	_, err := io.Copy(&payload, r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println("handle action", payload.String())
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
