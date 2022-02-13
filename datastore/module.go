package datastore

import (
	"github.com/syllabix/oncall/datastore/db"
	"github.com/syllabix/oncall/datastore/schedule"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	db.SetupDatabase,
	schedule.NewStore,
)
