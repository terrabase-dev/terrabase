package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/terrabase-dev/terrabase/internal/models"
	"github.com/uptrace/bun"
)

type CredentialRepo struct {
	db *bun.DB
}

func NewCredentialRepo(db *bun.DB) *CredentialRepo {
	return &CredentialRepo{db: db}
}

func (r *CredentialRepo) UpsertPassword(ctx context.Context, userID string, passwordHash string, algorithm string) error {
	now := time.Now().UTC()

	cred := &models.UserCredential{
		UserID:       userID,
		PasswordHash: passwordHash,
		Algorithm:    algorithm,
		PasswordSet:  now,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	_, err := r.db.NewInsert().
		Model(cred).
		On("CONFLICT (user_id) DO UPDATE").
		Set("password_hash = EXCLUDED.password_hash").
		Set("algorithm = EXCLUDED.algorithm").
		Set("password_set = EXCLUDED.password_set").
		Set("updated_at = EXCLUDED.updated_at").
		Exec(ctx)

	return err
}

func (r *CredentialRepo) GetByUserID(ctx context.Context, userID string) (*models.UserCredential, error) {
	var cred models.UserCredential

	err := r.db.NewSelect().Model(&cred).Where("user_id = ?", userID).Scan(ctx)

	if err != nil {
		if isNotFound(err) {
			return nil, fmt.Errorf("credential: %w", ErrNotFound)
		}

		return nil, err
	}

	return &cred, nil
}
