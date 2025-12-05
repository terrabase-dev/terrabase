package models

import (
	"time"

	applicationv1 "github.com/terrabase-dev/terrabase/specs/terrabase/application/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Application struct {
	bun.BaseModel `bun:"table:application"`

	ID        string    `bun:",pk"`
	Name      string    `bun:",unique,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	TeamRef                        *Team                         `bun:"rel:belongs-to,join:team_id=id"`
	TeamApplicationAccessGrantsRef *[]TeamApplicationAccessGrant `bun:"rel:has-many,join:application_id=id"`
}

func ApplicationFromProto(application *applicationv1.Application) *Application {
	if application == nil {
		return &Application{}
	}

	return &Application{
		ID:   application.GetId(),
		Name: application.GetName(),
	}
}

func (a *Application) ToProto() *applicationv1.Application {
	return &applicationv1.Application{
		Id:        a.ID,
		Name:      a.Name,
		CreatedAt: timestamppb.New(a.CreatedAt.UTC()),
		UpdatedAt: timestamppb.New(a.UpdatedAt.UTC()),
	}
}

func (a *Application) SetUpdatedAt(updatedAt time.Time) {
	a.UpdatedAt = updatedAt
}

func (a *Application) ModelName() string {
	return "application"
}
