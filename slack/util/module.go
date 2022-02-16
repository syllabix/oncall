package util

import (
	"github.com/syllabix/oncall/slack/util/mention"
	"go.uber.org/fx"
)

// Module exposes dependencies useful for dealing with the Slack
// apis and data
var Module = fx.Provide(
	mention.NewParser,
)
