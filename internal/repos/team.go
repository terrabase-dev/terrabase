package repos

import (
	"context"

	"github.com/terrabase-dev/terrabase/internal/models"
	teamv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team/v1"
	"github.com/uptrace/bun"
)

type TeamRepo struct {
	db *bun.DB
}

func NewTeamRepo(db *bun.DB) *TeamRepo {
	return &TeamRepo{
		db: db,
	}
}

func (r *TeamRepo) Create(ctx context.Context, team *teamv1.Team) (*teamv1.Team, error) {
	model := models.TeamFromProto(team)

	if model.ID == "" {
		model.ID = uuidString()
	}

	return create(ctx, r.db, model)
}

func (r *TeamRepo) Get(ctx context.Context, id string) (*teamv1.Team, error) {
	model := new(models.Team)

	return get(ctx, r.db.NewSelect().Model(model), model, id)
}

func (r *TeamRepo) List(ctx context.Context, pageSize int32, pageToken string) ([]*teamv1.Team, string, error) {
	var models []*models.Team

	return paginate(ctx, r.db.NewSelect().Model(&models).Order("created_at DESC"), &models, pageSize, pageToken)
}

func (r *TeamRepo) Update(ctx context.Context, id string, name *string) (*teamv1.Team, error) {
	model := new(models.Team)

	if _, err := get(ctx, r.db.NewSelect().Model(model), model, id); err != nil {
		return nil, err
	}

	model.Name = *name

	return update(ctx, r.db, model, "name")
}

func (r *TeamRepo) Delete(ctx context.Context, id string) error {
	return delete(ctx, r.db, (*models.Team)(nil), id)
}
