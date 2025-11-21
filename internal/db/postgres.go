package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// OpenPostgres returns a Bun DB backed by the provided DSN.
func OpenPostgres(ctx context.Context, dsn string) (*bun.DB, error) {
	if dsn == "" {
		return nil, errors.New("DATABASE_URL/DSN is required")
	}

	connector := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	sqldb := sql.OpenDB(connector)

	if err := sqldb.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())
	return db, nil
}
