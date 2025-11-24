package models

import (
	"time"

	"github.com/uptrace/bun"
)

// Identity links a user to an external auth provider (GitHub/GitLab/OIDC/etc).
type Identity struct {
	bun.BaseModel `bun:"table:identity"`

	ID             string `bun:",pk"`
	UserID         string `bun:",notnull"`
	Provider       string `bun:",notnull"`
	ProviderUserID string `bun:",notnull"`
	Email          string
	Metadata       map[string]any `bun:",type:jsonb"`
	User           *User          `bun:"rel:belongs-to,join:user_id=id"`
	CreatedAt      time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
}
