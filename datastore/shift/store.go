package shift

import (
	"context"

	"github.com/syllabix/oncall/common/db"
	"github.com/syllabix/oncall/datastore/model"
)

type Store struct {
	db db.Postgres
}

func NewStore(db db.Postgres) (Store, error) {
	return Store{db}, nil
}

func (s Store) Update(ctx context.Context, shifts ...*model.Shift) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return failure(err)
	}

	_, err = tx.NamedExec(update, shifts)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return failure(err)
	}

	return nil
}

func (s Store) Delete(ctx context.Context, shift *model.Shift) error {
	_, err := shift.Delete(ctx, s.db)
	if err != nil {
		return failure(err)
	}
	return nil
}
