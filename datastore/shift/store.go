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
		err = tx.Shift.UpdateOneID(s.ID).
			SetSequenceID(s.SequenceID).
			SetScheduleID(s.ScheduleID).
			SetStatus(s.Status).
			SetUserID(s.UserID).
			Exec(ctx)
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
