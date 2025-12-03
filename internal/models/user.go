package models

import (
	"time"

	userv1 "github.com/terrabase-dev/terrabase/specs/terrabase/user/v1"
	userRolev1 "github.com/terrabase-dev/terrabase/specs/terrabase/user_role/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// User represents a human user account. We avoid the reserved "user" table name
// by using "user_account".
type User struct {
	bun.BaseModel `bun:"table:user_account"`

	ID           string `bun:",pk"`
	Name         string
	Email        string `bun:",unique"`
	DefaultRole  int32
	UserType     userv1.UserType
	OwnerUserID  string         `bun:",nullzero,on_delete:CASCADE"`
	Metadata     map[string]any `bun:",type:jsonb"`
	CreatedAt    time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time      `bun:",nullzero,notnull,default:current_timestamp"`
	OwnerUserRef *User          `bun:"rel:belongs-to,join:owner_user_id=id"`
}

func UserFromProto(user *userv1.User) *User {
	if user == nil {
		return &User{}
	}

	return &User{
		ID:          user.GetId(),
		Name:        user.GetName(),
		Email:       user.GetEmail(),
		DefaultRole: int32(user.GetDefaultRole()),
		UserType:    user.GetUserType(),
		OwnerUserID: user.GetOwnerUserId(),
		Metadata:    nil,
	}
}

func (u *User) ToProto() *userv1.User {
	res := &userv1.User{
		Id:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		DefaultRole: userRolev1.UserRole(u.DefaultRole),
		UserType:    userv1.UserType(u.UserType),
		CreatedAt:   timestamppb.New(u.CreatedAt.UTC()),
		UpdatedAt:   timestamppb.New(u.UpdatedAt.UTC()),
	}

	if u.OwnerUserID != "" {
		res.OwnerUserId = &u.OwnerUserID
	}

	return res
}
