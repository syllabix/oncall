package slack

import (
	"github.com/syllabix/oncall/service/slack/event"
	"go.uber.org/fx"
)

// Module contains all services that power handling
// interactions with Slack api calls
var Module = fx.Provide(
	event.NewParser,
	event.NewHandler,
)
