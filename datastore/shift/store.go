package shift

import (
	"context"

	"github.com/syllabix/oncall/datastore/entity"
)

type Store struct {
	db *entity.Client
}

func NewStore(db *entity.Client) Store {
	return Store{db}
}

func (s Store) Update(ctx context.Context, shifts ...*entity.Shift) error {
	tx, err := s.db.Tx(ctx)
	if err != nil {
		return failure(err)
	}

	for _, s := range shifts {
		_, err = tx.Shift.UpdateOne(s).Save(ctx)
		if err != nil {
			return failure(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return failure(err)
	}

	return nil
}

func (s Store) Delete(ctx context.Context, shift *entity.Shift) error {
	err := s.db.Shift.DeleteOne(shift).Exec(ctx)
	if err != nil {
		return failure(err)
	}
	return nil
}
