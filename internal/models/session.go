package models

import (
	"time"

	"github.com/uptrace/bun"
)

// Session tracks refresh tokens issued for interactive logins.
type Session struct {
	bun.BaseModel `bun:"table:session"`

	ID               string    `bun:",pk"`
	UserID           string    `bun:",notnull"`
	RefreshTokenHash string    `bun:",notnull"`
	ExpiresAt        time.Time `bun:",notnull"`
	LastUsedAt       time.Time `bun:",nullzero"`
	UserAgent        string
	IP               string
	Metadata         map[string]any `bun:",type:jsonb"`
	User             *User          `bun:"rel:belongs-to,join:user_id=id"`
	CreatedAt        time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt        time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
}
