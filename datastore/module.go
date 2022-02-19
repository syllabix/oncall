package datastore

import (
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/datastore/shift"
	"github.com/syllabix/oncall/datastore/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewStore,
	schedule.NewStore,
	shift.NewStore,
)
