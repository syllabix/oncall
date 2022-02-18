package service

import (
	"github.com/syllabix/oncall/service/schedule"
	"go.uber.org/fx"
)

// Module contains all services that power handling
// interactions with Slack api calls
var Module = fx.Options(

	schedule.Module,
)
