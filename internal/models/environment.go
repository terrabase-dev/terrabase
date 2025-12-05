package models

import (
	"time"

	environmentv1 "github.com/terrabase-dev/terrabase/specs/terrabase/environment/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Environment struct {
	bun.BaseModel `bun:"table:environment"`

	ID            string    `bun:",pk"`
	Name          string    `bun:",notnull"`
	ApplicationID string    `bun:",nullzero,notnull,on_delete:CASCADE"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	ApplicationRef *Application `bun:"rel:belongs-to,join:application_id=id"`
}

func EnvironmentFromProto(environment *environmentv1.Environment) *Environment {
	if environment == nil {
		return &Environment{}
	}

	return &Environment{
		ID:            environment.GetId(),
		Name:          environment.GetName(),
		ApplicationID: environment.GetApplicationId(),
	}
}

func (e *Environment) ToProto() *environmentv1.Environment {
	return &environmentv1.Environment{
		Id:            e.ID,
		Name:          e.Name,
		ApplicationId: e.ApplicationID,
		CreatedAt:     timestamppb.New(e.CreatedAt.UTC()),
		UpdatedAt:     timestamppb.New(e.UpdatedAt.UTC()),
	}
}

func (a *Environment) SetUpdatedAt(updatedAt time.Time) {
	a.UpdatedAt = updatedAt
}

func (a *Environment) ModelName() string {
	return "environment"
}
