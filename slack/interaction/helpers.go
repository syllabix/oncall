package interaction

import (
	"github.com/slack-go/slack"
)

func actionTypeFor(callback slack.InteractionCallback) string {

	switch callback.Type {
	case slack.InteractionTypeBlockActions:
		if len(callback.ActionCallback.BlockActions) > 0 {
			return callback.ActionCallback.BlockActions[0].ActionID
		}
	}

	return "unknown"
}
