package repos

import (
	"context"

	"github.com/terrabase-dev/terrabase/internal/models"
	applicationv1 "github.com/terrabase-dev/terrabase/specs/terrabase/application/v1"
	"github.com/uptrace/bun"
)

type ApplicationRepo struct {
	db bun.IDB
}

func NewApplicationRepo(db *bun.DB) *ApplicationRepo {
	return &ApplicationRepo{db: db}
}

func (r *ApplicationRepo) WithTx(tx bun.Tx) *ApplicationRepo {
	return &ApplicationRepo{db: tx}
}

func (r *ApplicationRepo) Create(ctx context.Context, application *applicationv1.Application) (*applicationv1.Application, error) {
	model := models.ApplicationFromProto(application)

	if model.ID == "" {
		model.ID = uuidString()
	}

	return create(ctx, r.db, model)
}

func (r *ApplicationRepo) Get(ctx context.Context, id string) (*applicationv1.Application, error) {
	model := new(models.Application)

	return get(ctx, r.db.NewSelect().Model(model), model, id)
}

func (r *ApplicationRepo) List(ctx context.Context, teamId string, pageSize int32, pageToken string) ([]*applicationv1.Application, string, error) {
	var applications []*models.Application

	// Ensure the team exists
	teamModel := new(models.Team)
	if _, err := get(ctx, r.db.NewSelect().Model(teamModel), teamModel, teamId); err != nil {
		return nil, "", err
	}

	query := r.db.NewSelect().
		Model(&applications).
		Relation("TeamApplicationAccessGrantsRef").
		Join(`JOIN "TeamApplication" ta ON ta.application_id = application.id`).
		Where("ta.team_id = ?", teamId).
		OrderExpr("application.updated_at DESC")

	return paginate(ctx, query, &applications, pageSize, pageToken)
}

func (r *ApplicationRepo) Update(ctx context.Context, id string, name string) (*applicationv1.Application, error) {
	model := new(models.Application)

	if _, err := get(ctx, r.db.NewSelect().Model(model), model, id); err != nil {
		return nil, err
	}

	model.Name = name

	return update(ctx, r.db, model, "name")
}

func (r *ApplicationRepo) Delete(ctx context.Context, id string) error {
	return delete(ctx, r.db, (*models.Application)(nil), id)
}
