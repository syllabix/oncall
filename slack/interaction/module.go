package interaction

import (
	"github.com/syllabix/oncall/slack/interaction/schedule"
	"go.uber.org/fx"
)

// Module provided all the dependencies that deal
// with Slack interaction handling
var Module = fx.Provide(
	NewHandler,
	schedule.NewCreator,
)
