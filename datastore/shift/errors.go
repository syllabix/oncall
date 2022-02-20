package shift

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("the requested resource was not found")
	ErrFatal    = errors.New("an unexpected error occurred while accessing the database")
	ErrConflict = errors.New("an entity has already been created for this resource type")
)

func mapErr(reason error) (err error) {
	var sqlErr *pq.Error
	switch {
	case errors.Is(reason, sql.ErrNoRows):
		err = ErrNotFound

	case errors.As(reason, &sqlErr):
		if sqlErr.Code.Name() == "unique_violation" {
			err = ErrConflict
			return
		}

		// if it is not an expected SQL error, then use default case
		fallthrough

	default:
		err = fmt.Errorf("%w: %v", ErrFatal, reason)
	}

	return errors.WithStack(err)
}

func failure(reason error) (err error) {
	return mapErr(reason)
}

func rollback(tx *sqlx.Tx, reason error) (err error) {
	err = tx.Rollback()
	if err != nil {
		return mapErr(err)
	}

	return mapErr(reason)
}
