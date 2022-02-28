package controller

import (
	"github.com/syllabix/oncall/api/controller/command"
	"github.com/syllabix/oncall/api/controller/health"
	"github.com/syllabix/oncall/api/controller/interaction"
	"github.com/syllabix/oncall/api/controller/slack"
	"go.uber.org/fx"
)

// Module contains all the of the On Call
// API controllers
var Module = fx.Provide(
	slack.NewController,
	command.NewController,
	health.NewController,
	interaction.NewController,
)
