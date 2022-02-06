package client

import (
	"github.com/syllabix/oncall/client/slack"
	"go.uber.org/fx"
)

// Module contains all of the dependencies required
// to interface with 3rd party clients
var Module = fx.Provide(
	slack.NewClient,
)
