package models

import (
	"errors"
	"fmt"
	"time"

	workspacev1 "github.com/terrabase-dev/terrabase/specs/terrabase/workspace/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Workspace struct {
	bun.BaseModel `bun:"table:workspace"`

	ID            string                  `bun:",pk"`
	Name          string                  `bun:",unique,notnull"`
	BackendType   workspacev1.BackendType `bun:",notnull"`
	EnvironmentID string                  `bun:",nullzero,on_delete:CASCADE"`
	TeamID        string                  `bun:",nullzero,on_delete:RESTRICT"`
	CreatedAt     time.Time               `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time               `bun:",nullzero,notnull,default:current_timestamp"`

	EnvironmentRef     *Environment     `bun:"rel:belongs-to;join:environment_id:id"`
	TeamRef            *Team            `bun:"rel:belongs-to;join:team_id:id"`
	S3BackendConfigRef *S3BackendConfig `bun:"rel:has-one,join:id=workspace_id"`
}

func WorkspaceFromProto(workspace *workspacev1.Workspace) (*Workspace, error) {
	if workspace == nil {
		return &Workspace{}, nil
	}

	res := &Workspace{
		ID:          workspace.GetId(),
		Name:        workspace.GetName(),
		BackendType: workspace.GetBackendType(),
	}

	switch o := workspace.Owner.(type) {
	case *workspacev1.Workspace_EnvironmentId:
		res.EnvironmentID = o.EnvironmentId
	case *workspacev1.Workspace_TeamId:
		res.TeamID = o.TeamId
	default:
		return nil, fmt.Errorf("unknown workspace owner type %T", o)
	}

	return res, nil
}

func (w *Workspace) ToProto() (*workspacev1.Workspace, error) {
	res := &workspacev1.Workspace{
		Id:          w.ID,
		Name:        w.Name,
		BackendType: w.BackendType,
		CreatedAt:   timestamppb.New(w.CreatedAt.UTC()),
		UpdatedAt:   timestamppb.New(w.UpdatedAt.UTC()),
	}

	if w.EnvironmentID != "" && w.TeamID != "" {
		return nil, errors.New("cannot provide both environment_id and team_id")
	}

	switch {
	case w.EnvironmentID != "":
		res.Owner = &workspacev1.Workspace_EnvironmentId{EnvironmentId: w.EnvironmentID}
	case w.TeamID != "":
		res.Owner = &workspacev1.Workspace_TeamId{TeamId: w.TeamID}
	default:
		return nil, errors.New("must provide one of environment_id or team_id")
	}

	return res, nil
}
