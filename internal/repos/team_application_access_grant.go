package repos

import (
	"context"
	"errors"

	"github.com/terrabase-dev/terrabase/internal/models"
	teamAccessTypev1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_access_type/v1"
	teamApplicationAccessGrantv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_application_access_grant/v1"
	"github.com/uptrace/bun"
)

type TeamApplicationAccessGrantRepo struct {
	db *bun.DB
}

func NewTeamApplicationRepo(db *bun.DB) *TeamApplicationAccessGrantRepo {
	return &TeamApplicationAccessGrantRepo{
		db: db,
	}
}

func (r *TeamApplicationAccessGrantRepo) Create(ctx context.Context, teamApplicationAccessGrant *teamApplicationAccessGrantv1.TeamApplicationAccessGrant) (*teamApplicationAccessGrantv1.TeamApplicationAccessGrant, error) {
	// Ensure the team exists
	teamModel := new(models.Team)
	teamId := teamApplicationAccessGrant.GetTeamId()
	if _, err := get(ctx, r.db.NewSelect().Model(teamModel), teamModel, teamId); err != nil {
		return nil, err
	}

	// Ensure the application exists
	applicationModel := new(models.Application)
	applicationId := teamApplicationAccessGrant.GetApplicationId()
	if _, err := get(ctx, r.db.NewSelect().Model(applicationModel), applicationModel, applicationId); err != nil {
		return nil, err
	}

	model := models.TeamApplicationAccessGrantFromProto(teamApplicationAccessGrant)

	if model.ID == "" {
		model.ID = uuidString()
	}

	return create(ctx, r.db, model)
}

func (r *TeamApplicationAccessGrantRepo) Get(ctx context.Context, id string) (*teamApplicationAccessGrantv1.TeamApplicationAccessGrant, error) {
	model := new(models.TeamApplicationAccessGrant)

	return get(ctx, r.db.NewSelect().Model(model), model, id)
}

func (r *TeamApplicationAccessGrantRepo) List(ctx context.Context, teamId *string, applicationId *string, pageSize int32, pageToken string) ([]*teamApplicationAccessGrantv1.TeamApplicationAccessGrant, string, error) {
	var teamApplicationAccessGrants []*models.TeamApplicationAccessGrant

	query := r.db.NewSelect().Model(&teamApplicationAccessGrants).Order("updated_at DESC")

	if teamId != nil && applicationId != nil {
		return nil, "", errors.New("cannot provide both team_id and application_id")
	}

	if teamId != nil {
		// Ensure the team exists
		teamModel := new(models.Team)
		if _, err := get(ctx, r.db.NewSelect().Model(teamModel), teamModel, *teamId); err != nil {
			return nil, "", err
		}

		query = query.Where("team_id = ?", teamId)
	} else if applicationId != nil {
		// Ensure the application exists
		applicationModel := new(models.Application)
		if _, err := get(ctx, r.db.NewSelect().Model(applicationModel), applicationModel, *applicationId); err != nil {
			return nil, "", err
		}

		query = query.Where("application_id = ?", applicationId)
	}

	return paginate(ctx, query, &teamApplicationAccessGrants, pageSize, pageToken)
}

func (r *TeamApplicationAccessGrantRepo) Update(ctx context.Context, id string, accessType teamAccessTypev1.TeamAccessType) (*teamApplicationAccessGrantv1.TeamApplicationAccessGrant, error) {
	model := new(models.TeamApplicationAccessGrant)

	if _, err := get(ctx, r.db.NewSelect().Model(model), model, id); err != nil {
		return nil, err
	}

	model.AccessType = accessType

	return update(ctx, r.db, model, "access_type")
}

func (r *TeamApplicationAccessGrantRepo) Delete(ctx context.Context, teamId string, applicationId string) error {
	model := new(models.TeamApplicationAccessGrant)

	if _, err := r.db.NewSelect().Model(model).Where("team_id = ?", teamId).Where("application_id = ?", applicationId).Exec(ctx); err != nil {
		return err
	}

	return r.DeleteById(ctx, model.ID)
}

func (r *TeamApplicationAccessGrantRepo) DeleteById(ctx context.Context, id string) error {
	return delete(ctx, r.db, (*models.TeamApplicationAccessGrant)(nil), id)
}
