package datastore

import (
	"github.com/syllabix/oncall/datastore/db"
	"github.com/syllabix/oncall/datastore/schedule"
	"github.com/syllabix/oncall/datastore/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	db.SetupDatabase,

	user.NewStore,
	schedule.NewStore,
)
