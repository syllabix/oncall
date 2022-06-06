package database

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Open new connection
func Open(config Settings) (*entsql.Driver, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, fmt.Errorf("could not setup pgx db driver: %w", err)
	}

	db.SetMaxOpenConns(config.MaxConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetConnMaxLifetime(config.MaxConnLifetime)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to establish database connection: %w", err)
	}

	return entsql.OpenDB(dialect.Postgres, db), nil
}
