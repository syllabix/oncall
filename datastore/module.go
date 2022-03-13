package datastore

import (
	"github.com/syllabix/oncall/datastore/entity"
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/datastore/shift"
	"github.com/syllabix/oncall/datastore/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	entity.Module,
	fx.Provide(
		user.NewStore,
		schedule.NewStore,
		shift.NewStore,
	),
)
