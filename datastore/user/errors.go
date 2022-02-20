package user

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	ErrNotFound     = errors.New("the requested resource was not found")
	ErrFatal        = errors.New("an unexpected error occurred while accessing the database")
	ErrAlreadyInUse = errors.New("an entity has already been created for this resource type")
)

func mapErr(reason error) (err error) {
	var sqlErr *pq.Error
	switch {
	case errors.Is(reason, sql.ErrNoRows):
		err = ErrNotFound

	case errors.As(reason, &sqlErr):
		if sqlErr.Code.Name() == "unique_violation" {
			err = ErrAlreadyInUse
			return
		}

		// if it is not an expected SQL error, then use default case
		fallthrough

	default:
		err = fmt.Errorf("%w: %v", ErrFatal, reason)
	}

	return errors.WithStack(err)
}

func failure[T any](reason error) (res T, err error) {
	return res, mapErr(reason)
}

func rollback[T any](tx *sqlx.Tx, reason error) (res T, err error) {
	err = tx.Rollback()
	if err != nil {
		return res, mapErr(err)
	}

	return res, mapErr(reason)
}
