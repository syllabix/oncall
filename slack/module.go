package slack

import (
	"github.com/syllabix/oncall/slack/command"
	"github.com/syllabix/oncall/slack/interaction"
	"github.com/syllabix/oncall/slack/notifications"
	"github.com/syllabix/oncall/slack/util"
	"go.uber.org/fx"
)

// Module contains all of the services that interpret and
// handle Slack api messages, commands, interactions, events, etc
var Module = fx.Options(
	interaction.Module,
	util.Module,
	command.Module,
	notifications.Module,
)
