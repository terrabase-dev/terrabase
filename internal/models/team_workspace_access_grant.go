package models

import (
	"time"

	teamAccessTypev1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_access_type/v1"
	teamWorkspaceAccessGrantv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_workspace_access_grant/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TeamWorkspaceAccessGrant struct {
	bun.BaseModel `bun:"table:TeamWorkspace"`

	ID          string                          `bun:",pk"`
	TeamID      string                          `bun:",nullzero,notnull,unique:group:team_workspace_id,on_delete:CASCADE"`
	WorkspaceID string                          `bun:",nullzero,notnull,unique:group:team_workspace_id,on_delete:CASCADE"`
	AccessType  teamAccessTypev1.TeamAccessType `bun:",nullzero,notnull"`
	CreatedAt   time.Time                       `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time                       `bun:",nullzero,notnull,default:current_timestamp"`

	TeamRef      *Team      `bun:"rel:belongs-to,join:team_id=id"`
	WorkspaceRef *Workspace `bun:"rel:belongs-to,join:workspace_id=id"`
}

func TeamWorkspaceAccessGrantFromProto(teamWorkspaceAccessGrant *teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant) *TeamWorkspaceAccessGrant {
	if teamWorkspaceAccessGrant == nil {
		return &TeamWorkspaceAccessGrant{}
	}

	return &TeamWorkspaceAccessGrant{
		ID:          teamWorkspaceAccessGrant.GetId(),
		TeamID:      teamWorkspaceAccessGrant.GetTeamId(),
		WorkspaceID: teamWorkspaceAccessGrant.GetWorkspaceId(),
		AccessType:  teamWorkspaceAccessGrant.GetAccessType(),
	}
}

func (t *TeamWorkspaceAccessGrant) ToProto() *teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant {
	return &teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant{
		Id:          t.ID,
		TeamId:      t.TeamID,
		WorkspaceId: t.WorkspaceID,
		AccessType:  t.AccessType,
		CreatedAt:   timestamppb.New(t.CreatedAt.UTC()),
		UpdatedAt:   timestamppb.New(t.UpdatedAt.UTC()),
	}
}

func (t *TeamWorkspaceAccessGrant) SetUpdatedAt(updatedAt time.Time) {
	t.UpdatedAt = updatedAt
}

func (t *TeamWorkspaceAccessGrant) ModelName() string {
	return "team_workspace"
}
