package db

import (
	"context"
	"fmt"
	"log"

	"github.com/terrabase-dev/terrabase/internal/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var migrations = migrate.NewMigrations()

func init() {
	migrations.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.NewCreateTable().
				IfNotExists().
				Model((*models.Organization)(nil)).
				WithForeignKeys().
				Exec(ctx)
			return err
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.NewDropTable().
				IfExists().
				Model((*models.Organization)(nil)).
				Exec(ctx)
			return err
		},
	)
}

// RunMigrations applies all pending migrations.
func RunMigrations(ctx context.Context, db *bun.DB, logger *log.Logger) error {
	migrator := migrate.NewMigrator(db, migrations)

	if err := migrator.Init(ctx); err != nil {
		return fmt.Errorf("init migrations: %w", err)
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}

	if logger != nil {
		if group.IsZero() {
			logger.Printf("migrations: no changes")
		} else {
			logger.Printf("migrations: applied %s", group)
		}
	}

	return nil
}
