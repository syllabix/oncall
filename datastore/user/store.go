package user

import (
	"context"

	"github.com/syllabix/oncall/datastore/entity"
	"github.com/syllabix/oncall/datastore/entity/user"
)

type Store struct {
	db *entity.UserClient
}

func NewStore(db *entity.Client) Store {
	return Store{db.User}
}

func (s Store) GetByID(ctx context.Context, id int) (entity.User, error) {
	user, err := s.db.Get(ctx, id)
	if err != nil {
		return failure[entity.User](err)
	}
	return *user, nil
}

func (s Store) GetBySlackID(ctx context.Context, id string) (entity.User, error) {
	user, err := s.db.Query().
		Where(user.SlackIDEQ(id)).Only(ctx)
	if err != nil {
		return failure[entity.User](err)
	}
	return *user, nil
}

func (s Store) ListByID(ctx context.Context, ids ...int) ([]*entity.User, error) {
	users, err := s.db.Query().
		Where(user.IDIn(ids...)).All(ctx)
	if err != nil {
		return failure[[]*entity.User](err)
	}
	return users, nil
}

func (s Store) ListBySlackID(ctx context.Context, ids ...string) ([]*entity.User, error) {
	users, err := s.db.Query().
		Where(user.SlackIDIn(ids...)).All(ctx)
	if err != nil {
		return failure[[]*entity.User](err)
	}
	return users, nil
}
