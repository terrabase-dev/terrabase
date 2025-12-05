package repos

import (
	"database/sql"
	"errors"

	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	ErrNotFound          = errors.New("not found")
	ErrAlreadyExists     = errors.New("already exists")
	ErrNoUpdatesProvided = errors.New("no updates provided")
)

func isUniqueViolation(err error) bool {
	var pgErr *pgdriver.Error

	if errors.As(err, &pgErr) {
		// 23505: unique_violation
		return pgErr.Field('C') == "23505"
	}

	return false
}

func isNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
