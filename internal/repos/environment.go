package repos

import (
	"context"

	"github.com/terrabase-dev/terrabase/internal/models"
	environmentv1 "github.com/terrabase-dev/terrabase/specs/terrabase/environment/v1"
	"github.com/uptrace/bun"
)

type EnvironmentRepo struct {
	db *bun.DB
}

func NewEnvironmentRepo(db *bun.DB) *EnvironmentRepo {
	return &EnvironmentRepo{
		db: db,
	}
}

func (r *EnvironmentRepo) Create(ctx context.Context, environment *environmentv1.Environment) (*environmentv1.Environment, error) {
	// Ensure the application exists
	applicationModel := new(models.Application)
	applicationId := environment.GetApplicationId()
	if _, err := get(ctx, r.db.NewSelect().Model(applicationModel), applicationModel, applicationId); err != nil {
		return nil, err
	}

	model := models.EnvironmentFromProto(environment)

	if model.ID == "" {
		model.ID = uuidString()
	}

	return create(ctx, r.db, model)
}

func (r *EnvironmentRepo) Get(ctx context.Context, id string) (*environmentv1.Environment, error) {
	model := new(models.Environment)

	return get(ctx, r.db.NewSelect().Model(model), model, id)
}

func (r *EnvironmentRepo) List(ctx context.Context, applicationId string) ([]*environmentv1.Environment, error) {
	var models []*models.Environment

	err := r.db.NewSelect().Model(&models).Where("application_id = ?", applicationId).Order("updated_at DESC").Scan(ctx)
	if err != nil {
		return nil, err
	}

	environments := make([]*environmentv1.Environment, 0, len(models))

	for i := range models {
		environments = append(environments, models[i].ToProto())
	}

	return environments, nil
}

func (r *EnvironmentRepo) Update(ctx context.Context, id string, name *string) (*environmentv1.Environment, error) {
	model := new(models.Environment)

	if _, err := get(ctx, r.db.NewSelect().Model(model), model, id); err != nil {
		return nil, err
	}

	model.Name = *name

	return update(ctx, r.db, model, "name")
}

func (r *EnvironmentRepo) Delete(ctx context.Context, id string) error {
	return delete(ctx, r.db, (*models.Environment)(nil), id)
}
