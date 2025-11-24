package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/terrabase-dev/terrabase/internal/models"
	"github.com/uptrace/bun"
)

type SessionRepo struct {
	db *bun.DB
}

func NewSessionRepo(db *bun.DB) *SessionRepo {
	return &SessionRepo{db: db}
}

func (r *SessionRepo) Create(ctx context.Context, sessionID string, userID string, refreshTokenHash string, expiresAt time.Time, userAgent string, ip string, metadata map[string]any) (*models.Session, error) {
	now := time.Now().UTC()

	session := &models.Session{
		ID:               sessionIDOrNew(sessionID),
		UserID:           userID,
		RefreshTokenHash: refreshTokenHash,
		ExpiresAt:        expiresAt.UTC(),
		LastUsedAt:       now,
		UserAgent:        userAgent,
		IP:               ip,
		Metadata:         metadata,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	_, err := r.db.NewInsert().Model(session).Exec(ctx)

	return session, err
}

func (r *SessionRepo) GetByRefreshHash(ctx context.Context, refreshHash string) (*models.Session, error) {
	var session models.Session

	err := r.db.NewSelect().
		Model(&session).
		Where("refresh_token_hash = ?", refreshHash).
		Where("expires_at > NOW()").
		Scan(ctx)

	if err != nil {
		if isNotFound(err) {
			return nil, fmt.Errorf("session: %w", ErrNotFound)
		}

		return nil, err
	}

	return &session, nil
}

func (r *SessionRepo) DeleteByID(ctx context.Context, id string) error {
	res, err := r.db.NewDelete().Model((*models.Session)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	if rowCount(res) == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *SessionRepo) TouchLastUsed(ctx context.Context, id string, when time.Time) error {
	_, err := r.db.NewUpdate().
		Model((*models.Session)(nil)).
		Set("last_used_at = ?", when).
		Set("updated_at = ?", when).
		Where("id = ?", id).
		Exec(ctx)

	return err
}

func (r *SessionRepo) ListByUser(ctx context.Context, userID string) ([]models.Session, error) {
	var sessions []models.Session

	err := r.db.NewSelect().
		Model(&sessions).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Scan(ctx)

	return sessions, err
}

func uuidString() string {
	return uuid.NewString()
}

func sessionIDOrNew(id string) string {
	if id != "" {
		return id
	}

	return uuidString()
}
