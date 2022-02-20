package user

import (
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

func (s Store) GetByID(id int) (model.User, error) {
	var user model.User
	err := s.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return model.User{}, nil
	}
	return user, nil
}

func (s Store) ListByID(ids ...int) ([]model.User, error) {
	query, args, err := sqlx.In("SELECT * FROM users WHERE id IN (?);", ids)
	if err != nil {
		return nil, err
	}
	query = s.db.Rebind(query)

	var users []model.User
	err = s.db.Select(&users, query, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}
