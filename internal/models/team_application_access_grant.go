package models

import (
	"time"

	teamAccessTypev1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_access_type/v1"
	teamApplicationAccessGrantv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_application_access_grant/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TeamApplicationAccessGrant struct {
	bun.BaseModel `bun:"table:TeamApplication"`

	ID            string                          `bun:",pk"`
	TeamID        string                          `bun:",nullzero,notnull,unique:group:team_application_id,on_delete:CASCADE"`
	ApplicationID string                          `bun:",nullzero,notnull,unique:group:team_application_id,on_delete:CASCADE"`
	AccessType    teamAccessTypev1.TeamAccessType `bun:",nullzero,notnull"`
	CreatedAt     time.Time                       `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time                       `bun:",nullzero,notnull,default:current_timestamp"`

	TeamRef        *Team        `bun:"rel:belongs-to,join:team_id=id"`
	ApplicationRef *Application `bun:"rel:belongs-to,join:application_id=id"`
}

func TeamApplicationAccessGrantFromProto(teamApplicationAccessGrant *teamApplicationAccessGrantv1.TeamApplicationAccessGrant) *TeamApplicationAccessGrant {
	if teamApplicationAccessGrant == nil {
		return &TeamApplicationAccessGrant{}
	}

	return &TeamApplicationAccessGrant{
		ID:            teamApplicationAccessGrant.GetId(),
		TeamID:        teamApplicationAccessGrant.GetTeamId(),
		ApplicationID: teamApplicationAccessGrant.GetApplicationId(),
		AccessType:    teamApplicationAccessGrant.GetAccessType(),
	}
}

func (t *TeamApplicationAccessGrant) ToProto() (*teamApplicationAccessGrantv1.TeamApplicationAccessGrant, error) {
	return &teamApplicationAccessGrantv1.TeamApplicationAccessGrant{
		Id:            t.ID,
		TeamId:        t.TeamID,
		ApplicationId: t.ApplicationID,
		AccessType:    t.AccessType,
		CreatedAt:     timestamppb.New(t.CreatedAt.UTC()),
		UpdatedAt:     timestamppb.New(t.UpdatedAt.UTC()),
	}, nil
}

func (t *TeamApplicationAccessGrant) SetUpdatedAt(updatedAt time.Time) {
	t.UpdatedAt = updatedAt
}

func (t *TeamApplicationAccessGrant) ModelName() string {
	return "team_application"
}
