package command

import (
	"github.com/syllabix/oncall/slack/command/add"
	"go.uber.org/fx"
)

// Module exposes dependencies useful dealing with schedule
// related command processing
var Module = fx.Provide(
	NewHandler,
	add.NewHandler,
)
