package service

import (
	"github.com/syllabix/oncall/service/schedule"
	"github.com/syllabix/oncall/service/shift"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	shift.NewManager,
	schedule.NewManager,
)
