package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/syllabix/oncall/datastore/db"
	"github.com/syllabix/oncall/datastore/model"
)

type Store struct {
	db db.Database
}

func NewStore(db db.Database) Store {
	return Store{db}
}

func (s Store) GetByID(userID string) (model.User, error) {
	var user model.User
	err := s.db.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		return model.User{}, nil
	}
	return user, nil
}

func (s Store) GetAll(userIDs ...string) ([]model.User, error) {
	query, args, err := sqlx.In("SELECT * FROM users WHERE id IN (?);", userIDs)
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
