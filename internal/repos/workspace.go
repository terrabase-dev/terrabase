package repos

import (
	"context"
	"errors"

	"github.com/terrabase-dev/terrabase/internal/models"
	workspacev1 "github.com/terrabase-dev/terrabase/specs/terrabase/workspace/v1"
	"github.com/uptrace/bun"
)

type WorkspaceRepo struct {
	db bun.IDB
}

func NewWorkspaceRepo(db *bun.DB) *WorkspaceRepo {
	return &WorkspaceRepo{db: db}
}

func (r *WorkspaceRepo) WithTx(tx bun.Tx) *WorkspaceRepo {
	return &WorkspaceRepo{db: tx}
}

func (r *WorkspaceRepo) Create(ctx context.Context, workspace *workspacev1.Workspace) (*workspacev1.Workspace, error) {
	model := models.WorkspaceFromProto(workspace)

	if model.ID == "" {
		model.ID = uuidString()
	}

	return create(ctx, r.db, model)
}

func (r *WorkspaceRepo) Get(ctx context.Context, id string) (*workspacev1.Workspace, error) {
	model := new(models.Workspace)

	return get(ctx, r.db.NewSelect().Model(model), model, id)
}

func (r *WorkspaceRepo) List(ctx context.Context, teamId *string, applicationId *string, pageSize int32, pageToken string) ([]*workspacev1.Workspace, string, error) {
	var workspaces []*models.Workspace

	query := r.db.NewSelect().Model(&workspaces).OrderExpr("workspace.updated_at DESC")

	if teamId != nil && applicationId != nil {
		return nil, "", errors.New("cannot provide both team_id and application_id")
	}

	if teamId != nil {
		// Ensure the team exists
		teamModel := new(models.Team)
		if _, err := get(ctx, r.db.NewSelect().Model(teamModel), teamModel, *teamId); err != nil {
			return nil, "", err
		}

		query = query.
			Relation("TeamWorkspaceAccessGrantsRef").
			Join(`JOIN "TeamWorkspace" tw on tw.workspace_id = workspace.id`).
			Where("tw.team_id = ?", teamId)
	} else if applicationId != nil {
		// Ensure the application exists
		applicationModel := new(models.Application)
		if _, err := get(ctx, r.db.NewSelect().Model(applicationModel), applicationModel, *applicationId); err != nil {
			return nil, "", err
		}

		query = query.
			Relation("EnvironmentRef").
			Join(`JOIN "Environment" e on e.id = workspace.environment_id`).
			Where("e.application_id = ?", applicationId)
	}

	return paginate(ctx, query, &workspaces, pageSize, pageToken)
}

func (r *WorkspaceRepo) Update(ctx context.Context, id string, name *string, backendType *workspacev1.BackendType, environmentId *string) (*workspacev1.Workspace, error) {
	model := new(models.Workspace)

	if _, err := get(ctx, r.db.NewSelect().Model(model), model, id); err != nil {
		return nil, err
	}

	model.Name = *name
	model.BackendType = *backendType
	model.EnvironmentID = *environmentId

	return update(ctx, r.db, model, "name", "backend_type", "environment_id")
}

func (r *WorkspaceRepo) Delete(ctx context.Context, id string) error {
	return delete(ctx, r.db, (*models.Workspace)(nil), id)
}
