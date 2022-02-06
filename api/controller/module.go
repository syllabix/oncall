package controller

import (
	"github.com/syllabix/oncall/api/controller/slack"
	"go.uber.org/fx"
)

// Module contains all the of the On Call
// API controllers
var Module = fx.Provide(
	slack.NewController,
)
