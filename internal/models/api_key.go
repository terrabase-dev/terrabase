package models

import (
	"time"

	authzv1 "github.com/terrabase-dev/terrabase/specs/terrabase/authz/v1"
	"github.com/uptrace/bun"
)

// APIKey represents a long-lived key that can belong to a user, bot, or service principal.
type APIKey struct {
	bun.BaseModel `bun:"table:api_key"`

	ID            string          `bun:",pk"`
	OwnerType     string          `bun:",notnull"` // user | bot | service
	OwnerID       string          `bun:",notnull"`
	Name          string          `bun:",notnull"`
	Prefix        string          `bun:",notnull,unique"`
	SecretHash    string          `bun:",notnull"`
	Scopes        []authzv1.Scope `bun:",array"`
	Metadata      map[string]any  `bun:",type:jsonb"`
	ExpiresAt     time.Time       `bun:",nullzero"`
	RevokedAt     time.Time       `bun:",nullzero"`
	RevokedReason string
	RotatedFrom   string
	LastUsedAt    time.Time `bun:",nullzero"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
