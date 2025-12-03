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
		// User table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.User)(nil)).WithForeignKeys().Exec(ctx); err != nil {
			return err
		}

		// User Credential table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.UserCredential)(nil)).ForeignKey(`("user_id") REFERENCES "user_account" ("id") ON DELETE CASCADE`).Exec(ctx); err != nil {
			return err
		}

		// Identity table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.Identity)(nil)).ForeignKey(`("user_id") REFERENCES "user_account" ("id") ON DELETE CASCADE`).Exec(ctx); err != nil {
			return err
		}

		// Session table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.Session)(nil)).ForeignKey(`("user_id") REFERENCES "user_account" ("id") ON DELETE CASCADE`).Exec(ctx); err != nil {
			return err
		}

		// API Key table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.APIKey)(nil)).ForeignKey(`("owner_id") REFERENCES "user_account" ("id") ON DELETE CASCADE`).Exec(ctx); err != nil {
			return err
		}

		// Organization table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.Organization)(nil)).WithForeignKeys().Exec(ctx); err != nil {
			return err
		}

		// Team table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.Team)(nil)).ForeignKey(`("organization_id") REFERENCES "organization" ("id") ON DELETE CASCADE`).Exec(ctx); err != nil {
			return err
		}

		// Application table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.Application)(nil)).ForeignKey(`("team_id") REFERENCES "team" ("id") ON DELETE RESTRICT`).Exec(ctx); err != nil {
			return err
		}

		// Environment table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.Environment)(nil)).ForeignKey(`("application_id") REFERENCES "application" ("id") ON DELETE CASCADE`).Exec(ctx); err != nil {
			return err
		}

		// Workspace table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.Workspace)(nil)).ForeignKey(`("environment_id") REFERENCES "environment" ("id") ON DELETE CASCADE`).ForeignKey(`("team_id") REFERENCES "team" ("id") ON DELETE RESTRICT`).Exec(ctx); err != nil {
			return err
		}

		// S3 Backend Config table
		if _, err := db.NewCreateTable().IfNotExists().Model((*models.S3BackendConfig)(nil)).ForeignKey(`("workspace_id") REFERENCES "workspace" ("id") ON DELETE CASCADE`).Exec(ctx); err != nil {
			return err
		}

		return nil
	}

	down := func(ctx context.Context, db *bun.DB) error {
		tables := []any{
			(*models.S3BackendConfig)(nil),
			(*models.Workspace)(nil),
			(*models.Environment)(nil),
			(*models.Application)(nil),
			(*models.Team)(nil),
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
