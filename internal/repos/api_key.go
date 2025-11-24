package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/terrabase-dev/terrabase/internal/models"
	"github.com/uptrace/bun"
)

type APIKeyRepo struct {
	db *bun.DB
}

func NewAPIKeyRepo(db *bun.DB) *APIKeyRepo {
	return &APIKeyRepo{db: db}
}

func (r *APIKeyRepo) Create(ctx context.Context, key *models.APIKey) (*models.APIKey, error) {
	now := time.Now().UTC()

	key.CreatedAt = now
	key.UpdatedAt = now

	_, err := r.db.NewInsert().Model(key).Exec(ctx)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("api_key: %w", ErrAlreadyExists)
		}

		return nil, err
	}

	return key, nil
}

func (r *APIKeyRepo) GetByID(ctx context.Context, id string) (*models.APIKey, error) {
	var key models.APIKey

	err := r.db.NewSelect().Model(&key).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, fmt.Errorf("api_key: %w", ErrNotFound)
		}

		return nil, err
	}

	return &key, nil
}

func (r *APIKeyRepo) GetActiveByPrefix(ctx context.Context, prefix string) (*models.APIKey, error) {
	var key models.APIKey

	err := r.db.NewSelect().
		Model(&key).
		Where("prefix = ?", prefix).
		Where("revoked_at IS NULL").
		Where("expires_at IS NULL OR expires_at > NOW()").
		Scan(ctx)

	if err != nil {
		if isNotFound(err) {
			return nil, fmt.Errorf("api_key: %w", ErrNotFound)
		}

		return nil, err
	}

	return &key, nil
}

func (r *APIKeyRepo) TouchLastUsed(ctx context.Context, id string, when time.Time) error {
	_, err := r.db.NewUpdate().
		Model((*models.APIKey)(nil)).
		Set("last_used_at = ?", when).
		Set("updated_at = ?", when).
		Where("id = ?", id).
		Exec(ctx)

	return err
}

func (r *APIKeyRepo) ListByOwner(ctx context.Context, ownerType string, ownerID string) ([]*models.APIKey, error) {
	var keys []models.APIKey

	q := r.db.NewSelect().Model(&keys)

	if ownerType != "" {
		q = q.Where("owner_type = ?", ownerType)
	}

	if ownerID != "" {
		q = q.Where("owner_id = ?", ownerID)
	}

	err := q.OrderExpr("created_at DESC").Scan(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]*models.APIKey, 0, len(keys))

	for i := range keys {
		out = append(out, &keys[i])
	}

	return out, nil
}

func (r *APIKeyRepo) Revoke(ctx context.Context, id string, reason string) error {
	_, err := r.db.NewUpdate().
		Model((*models.APIKey)(nil)).
		Set("revoked_at = NOW()").
		Set("revoked_reason = ?", reason).
		Set("updated_at = NOW()").
		Where("id = ?", id).
		Exec(ctx)

	return err
}
