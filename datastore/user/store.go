package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/syllabix/oncall/common/db"
	"github.com/syllabix/oncall/datastore/model"
)

type Store struct {
	db db.Postgres
}

func NewStore(db db.Postgres) Store {
	return Store{db}
}

func (s Store) GetByID(ctx context.Context, id int) (model.User, error) {
	user, err := model.FindUser(ctx, s.db, id)
	if err != nil {
		return failure[model.User](err)
	}
	return *user, nil
}

func (s Store) GetBySlackID(ctx context.Context, id string) (model.User, error) {
	user, err := model.Users(
		model.UserWhere.SlackID.EQ(id),
	).One(ctx, s.db)
	if err != nil {
		return failure[model.User](err)
	}
	return *user, nil
}

func (s Store) ListByID(ids ...int) ([]model.User, error) {
	query, args, err := sqlx.In("SELECT * FROM users WHERE id IN (?);", ids)
	if err != nil {
		return failure[[]model.User](err)
	}
	query = s.db.Rebind(query)

	var users []model.User
	err = s.db.Select(&users, query, args...)
	if err != nil {
		return failure[[]model.User](err)
	}
	return users, nil
}

func (s Store) ListBySlackID(ids ...string) ([]model.User, error) {
	query, args, err := sqlx.In("SELECT * FROM users WHERE slack_id IN (?);", ids)
	if err != nil {
		return failure[[]model.User](err)
	}
	query = s.db.Rebind(query)

	var users []model.User
	err = s.db.Select(&users, query, args...)
	if err != nil {
		return failure[[]model.User](err)
	}
	return users, nil
}
