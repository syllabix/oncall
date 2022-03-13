package entity

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/syllabix/oncall/datastore/entity/database"
	"github.com/syllabix/oncall/datastore/entity/migrate"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func setup(lc fx.Lifecycle, driver *sql.Driver, log *zap.Logger) (*Client, error) {
	client := NewClient(Driver(driver))

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	err := client.Schema.
		Create(context.Background(),
			schema.WithAtlas(true),
			migrate.WithGlobalUniqueID(true),
			migrate.WithDropColumn(true),
			migrate.WithDropIndex(true),
		)
	if err != nil {
		return nil, err
	}

	log.Info("entity migrations run")

	return client, nil
}

var Module = fx.Provide(
	database.Open,
	setup,
)
