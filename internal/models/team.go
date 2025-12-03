package models

import (
	"time"

	teamv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Team struct {
	bun.BaseModel `bun:"table:team"`

	ID             string    `bun:",pk"`
	Name           string    `bun:",unique,notnull"`
	OrganizationID string    `bun:",notnull"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	Organization *Organization `bun:"rel:belongs-to,join:organization_id=id"`
}

func TeamFromProto(team *teamv1.Team) *Team {
	if team == nil {
		return &Team{}
	}

	return &Team{
		ID:             team.GetId(),
		Name:           team.GetName(),
		OrganizationID: team.GetOrganizationId(),
	}
}

func (t *Team) ToProto() (*teamv1.Team, error) {
	return &teamv1.Team{
		Id:             t.ID,
		Name:           t.Name,
		OrganizationId: t.OrganizationID,
		CreatedAt:      timestamppb.New(t.CreatedAt.UTC()),
		UpdatedAt:      timestamppb.New(t.UpdatedAt.UTC()),
	}, nil
}

func (t *Team) SetUpdatedAt(updatedAt time.Time) {
	t.UpdatedAt = updatedAt
}

func (t *Team) ModelName() string {
	return "team"
}
