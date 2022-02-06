package slack

import (
	"encoding/json"
	"net/http"

	"github.com/slack-go/slack/slackevents"
	"go.uber.org/zap"
)

func (ctrl *Controller) verifyURL(w http.ResponseWriter, event json.RawMessage) {
	var res slackevents.ChallengeResponse
	err := json.Unmarshal(event, &res)
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
}
