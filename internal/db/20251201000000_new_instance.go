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
	up := func(ctx context.Context, db *bun.DB) error {
		// Create user first so FKs have a target.
		if _, err := db.NewCreateTable().
			IfNotExists().
			Model((*models.User)(nil)).
			WithForeignKeys().
			Exec(ctx); err != nil {
			return err
		}

		// Dependent tables with explicit ON DELETE CASCADE where we want cascading cleanup.
		if _, err := db.NewCreateTable().
			IfNotExists().
			Model((*models.UserCredential)(nil)).
			ForeignKey(`("user_id") REFERENCES "user_account" ("id") ON DELETE CASCADE`).
			Exec(ctx); err != nil {
			return err
		}

		if _, err := db.NewCreateTable().
			IfNotExists().
			Model((*models.Identity)(nil)).
			ForeignKey(`("user_id") REFERENCES "user_account" ("id") ON DELETE CASCADE`).
			Exec(ctx); err != nil {
			return err
		}

		if _, err := db.NewCreateTable().
			IfNotExists().
			Model((*models.Session)(nil)).
			ForeignKey(`("user_id") REFERENCES "user_account" ("id") ON DELETE CASCADE`).
			Exec(ctx); err != nil {
			return err
		}

		if _, err := db.NewCreateTable().
			IfNotExists().
			Model((*models.APIKey)(nil)).
			ForeignKey(`("owner_id") REFERENCES "user_account" ("id") ON DELETE CASCADE`).
			Exec(ctx); err != nil {
			return err
		}

		if _, err := db.NewCreateTable().
			IfNotExists().
			Model((*models.Organization)(nil)).
			WithForeignKeys().
			Exec(ctx); err != nil {
			return err
		}

		if _, err := db.NewCreateTable().
			IfNotExists().
			Model((*models.Team)(nil)).
			ForeignKey(`("organization_id") REFERENCES "organization" ("id") ON DELETE CASCADE`).
			Exec(ctx); err != nil {
			return err
		}

		return nil
	}

	down := func(ctx context.Context, db *bun.DB) error {
		tables := []any{
			(*models.Organization)(nil),
			(*models.APIKey)(nil),
			(*models.Session)(nil),
			(*models.Identity)(nil),
			(*models.UserCredential)(nil),
			(*models.User)(nil),
		}

		for _, table := range tables {
			if _, err := db.NewDropTable().IfExists().Model(table).Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	}

	migrations.MustRegister(up, down)
}

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
