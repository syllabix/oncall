package db

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/fx"
	"go.uber.org/zap"

	// postgres driver required import
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrations embed.FS

const (
	postgres        = "postgres"
	migrationsTable = "migrations"
)

// Postgres is the primary database for the application.
type Postgres struct {
	*sqlx.DB
}

// Setup attempts configure and connect to the database
func Setup(lc fx.Lifecycle, config Settings, log *zap.Logger) (Postgres, error) {
	db, err := open(config)
	if err != nil {
		return Postgres{}, fmt.Errorf("unable to establish a connection with the db: %w", err)
	}

	err = ensure(db, log)
	if err != nil {
		return Postgres{}, fmt.Errorf("unable to establish a connection with the db: %w", err)
	}

	count, err := runMigrations(db)
	if err != nil {
		return Postgres{}, fmt.Errorf("database migrations failed: %w", err)
	}
	log.Info("database migrations completed ok", zap.Int("count", count))

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Info("closing down database connections...")
			err := db.Close()
			if err != nil {
				return fmt.Errorf("database connection failed to close properly: %w", err)
			}
			return nil
		},
	})

	return Postgres{db}, nil
}

func open(config Settings) (*sqlx.DB, error) {
	db, err := sqlx.Open(postgres, config.String())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetConnMaxLifetime(config.MaxConnLifetime)

	return db, nil
}

func ensure(db *sqlx.DB, log *zap.Logger) (err error) {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	retries := 0
	for retries < 10 {
		err = db.Ping()
		if err != nil {
			log.Warn("unable to establish postgres db connection: retrying...", zap.Error(err))
			<-ticker.C
			retries++
		} else {
			break
		}
	}

	return err
}

func runMigrations(db *sqlx.DB) (count int, err error) {
	fsys, err := fs.Sub(migrations, "migrations")
	if err != nil {
		return 0, err
	}

	source := migrate.HttpFileSystemMigrationSource{
		FileSystem: http.FS(fsys),
	}

	migrate.SetTable(migrationsTable)
	count, err = migrate.Exec(db.DB, postgres, source, migrate.Up)
	if err != nil {
		return 0, err
	}

	return count, err
}
