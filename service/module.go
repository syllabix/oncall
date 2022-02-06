package service

import (
	"github.com/syllabix/oncall/service/slack"
	"go.uber.org/fx"
)

// Module contains all services that power handling
// interactions with Slack api calls
var Module = fx.Options(
	slack.Module,
)
