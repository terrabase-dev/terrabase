package models

import (
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
	CreatedAt     time.Time               `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time               `bun:",nullzero,notnull,default:current_timestamp"`

	EnvironmentRef     *Environment     `bun:"rel:belongs-to;join:environment_id:id"`
	S3BackendConfigRef *S3BackendConfig `bun:"rel:has-one,join:id=workspace_id"`
}

func WorkspaceFromProto(workspace *workspacev1.Workspace) *Workspace {
	if workspace == nil {
		return &Workspace{}
	}

	res := &Workspace{
		ID:          workspace.GetId(),
		Name:        workspace.GetName(),
		BackendType: workspace.GetBackendType(),
	}

	if workspace.EnvironmentId != nil {
		res.EnvironmentID = *workspace.EnvironmentId
	}

	return res
}

func (w *Workspace) ToProto() *workspacev1.Workspace {
	res := &workspacev1.Workspace{
		Id:          w.ID,
		Name:        w.Name,
		BackendType: w.BackendType,
		CreatedAt:   timestamppb.New(w.CreatedAt.UTC()),
		UpdatedAt:   timestamppb.New(w.UpdatedAt.UTC()),
	}

	if w.EnvironmentID != "" {
		res.EnvironmentId = &w.EnvironmentID
	}

	return res
}

func (w *Workspace) SetUpdatedAt(updatedAt time.Time) {
	w.UpdatedAt = updatedAt
}

func (w *Workspace) ModelName() string {
	return "workspace"
}
