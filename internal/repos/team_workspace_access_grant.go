package repos

import (
	"context"

	"github.com/terrabase-dev/terrabase/internal/models"
	teamAccessTypev1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_access_type/v1"
	teamWorkspaceAccessGrantv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_workspace_access_grant/v1"
	"github.com/uptrace/bun"
)

type TeamWorkspaceAccessGrantRepo struct {
	db bun.IDB
}

func NewTeamWorkspaceRepo(db *bun.DB) *TeamWorkspaceAccessGrantRepo {
	return &TeamWorkspaceAccessGrantRepo{db: db}
}

func (r *TeamWorkspaceAccessGrantRepo) WithTx(tx bun.Tx) *TeamWorkspaceAccessGrantRepo {
	return &TeamWorkspaceAccessGrantRepo{db: tx}
}

func (r *TeamWorkspaceAccessGrantRepo) Create(ctx context.Context, teamWorkspaceAccessGrant *teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant) (*teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant, error) {
	// Ensure the team exists
	teamModel := new(models.Team)
	teamId := teamWorkspaceAccessGrant.GetTeamId()
	if _, err := get(ctx, r.db.NewSelect().Model(teamModel), teamModel, teamId); err != nil {
		return nil, err
	}

	// Ensure the workspace exists
	workspaceModel := new(models.Workspace)
	workspaceId := teamWorkspaceAccessGrant.GetWorkspaceId()
	if _, err := get(ctx, r.db.NewSelect().Model(workspaceModel), workspaceModel, workspaceId); err != nil {
		return nil, err
	}

	model := models.TeamWorkspaceAccessGrantFromProto(teamWorkspaceAccessGrant)

	if model.ID == "" {
		model.ID = uuidString()
	}

	return create(ctx, r.db, model)
}

func (r *TeamWorkspaceAccessGrantRepo) Get(ctx context.Context, id string) (*teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant, error) {
	model := new(models.TeamWorkspaceAccessGrant)

	return get(ctx, r.db.NewSelect().Model(model), model, id)
}

func (r *TeamWorkspaceAccessGrantRepo) List(ctx context.Context, teamId *string, workspaceId *string, pageSize int32, pageToken string) ([]*teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant, string, error) {
	var teamWorkspaceAccessGrants []*models.TeamWorkspaceAccessGrant

	query := r.db.NewSelect().Model(&teamWorkspaceAccessGrants).Order("updated_at DESC")

	if teamId != nil {
		// Ensure the team exists
		teamModel := new(models.Team)
		if _, err := get(ctx, r.db.NewSelect().Model(teamModel), teamModel, *teamId); err != nil {
			return nil, "", err
		}

		query = query.Where("team_id = ?", teamId)
	} else if workspaceId != nil {
		// Ensure the workspace exists
		workspaceModel := new(models.Workspace)
		if _, err := get(ctx, r.db.NewSelect().Model(workspaceModel), workspaceModel, *workspaceId); err != nil {
			return nil, "", err
		}

		query = query.Where("workspace_id = ?", workspaceId)
	}

	return paginate(ctx, query, &teamWorkspaceAccessGrants, pageSize, pageToken)
}

func (r *TeamWorkspaceAccessGrantRepo) Update(ctx context.Context, id string, accessType teamAccessTypev1.TeamAccessType) (*teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant, error) {
	model := new(models.TeamWorkspaceAccessGrant)

	if _, err := get(ctx, r.db.NewSelect().Model(model), model, id); err != nil {
		return nil, err
	}

	model.AccessType = accessType

	return update(ctx, r.db, model, "access_type")
}

func (r *TeamWorkspaceAccessGrantRepo) Delete(ctx context.Context, teamId string, workspaceId string) error {
	model := new(models.TeamWorkspaceAccessGrant)

	if _, err := r.db.NewSelect().Model(model).Where("team_id = ?", teamId).Where("workspace_id = ?", workspaceId).Exec(ctx); err != nil {
		return err
	}

	return r.DeleteById(ctx, model.ID)
}

func (r *TeamWorkspaceAccessGrantRepo) DeleteById(ctx context.Context, id string) error {
	return delete(ctx, r.db, (*models.TeamWorkspaceAccessGrant)(nil), id)
}
