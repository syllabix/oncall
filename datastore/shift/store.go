package shift

import (
	"context"

	"github.com/syllabix/oncall/common/db"
	"github.com/syllabix/oncall/datastore/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

	for _, shift := range shifts {
		_, err := shift.Update(ctx, tx, boil.Infer())
		if err != nil {
			return rollback(tx, err)
		}
	}

	return nil
}
