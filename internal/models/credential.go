package models

import (
	"time"

	"github.com/uptrace/bun"
)

// UserCredential stores password-based authentication data for a user.
type UserCredential struct {
	bun.BaseModel `bun:"table:user_credential"`

	UserID       string `bun:",pk"`
	User         *User  `bun:"rel:belongs-to,join:user_id=id"`
	PasswordHash string
	Algorithm    string
	PasswordSet  time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
