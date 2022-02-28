package db

import (
	"io"
	"log"
	"strings"
	"time"
)

type Tester struct {
	db Postgres
}

func NewTester() Tester {
	config := Settings{
		DBName:             "oncall",
		SSLMode:            "disable",
		User:               "admin",
		Password:           "p4ssw0rd",
		Host:               "localhost",
		Port:               "5432",
		MaxConnections:     1,
		MaxIdleConnections: 1,
		MaxConnLifetime:    time.Minute,
	}

	db, err := open(config)
	if err != nil {
		log.Fatalf("unable to establish a connection with the db: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to establish a connection with the db: %v", err)
	}

	_, err = runMigrations(db)
	if err != nil {
		log.Fatalf("database migrations failed: %v", err)
	}

	return Tester{Postgres{db}}
}

func (t Tester) GetDB() Postgres {
	return t.db
}

func (t Tester) Seed(r io.Reader) error {
	var sql strings.Builder
	_, err := io.Copy(&sql, r)
	if err != nil {
		return err
	}

	_, err = t.db.Exec(sql.String())
	if err != nil {
		return err
	}
	return nil
}

func (t Tester) Reset() error {
	_, err := t.db.Exec("DELETE from shifts")
	if err != nil {
		return err
	}

	_, err = t.db.Exec("DELETE from users")
	if err != nil {
		return err
	}

	_, err = t.db.Exec("DELETE from schedules")
	if err != nil {
		return err
	}

	return nil
}
