package repos

import (
	"context"

	"github.com/terrabase-dev/terrabase/internal/models"
	organizationv1 "github.com/terrabase-dev/terrabase/specs/terrabase/organization/v1"
	"github.com/uptrace/bun"
)

type OrganizationRepo struct {
	db *bun.DB
}

func NewOrganizationRepo(db *bun.DB) *OrganizationRepo {
	return &OrganizationRepo{
		db: db,
	}
}

func (r *OrganizationRepo) Create(ctx context.Context, org *organizationv1.Organization) (*organizationv1.Organization, error) {
	model := models.OrganizationFromProto(org)

	if model.ID == "" {
		model.ID = uuidString()
	}

	return create(ctx, r.db, model)
}

func (r *OrganizationRepo) Get(ctx context.Context, id string) (*organizationv1.Organization, error) {
	model := new(models.Organization)

	return get(ctx, r.db.NewSelect().Model(model), model, id)
}

func (r *OrganizationRepo) List(ctx context.Context, pageSize int32, pageToken string) ([]*organizationv1.Organization, string, error) {
	var models []*models.Organization

	return paginate(ctx, r.db.NewSelect().Model(&models).Order("created_at DESC"), &models, pageSize, pageToken)
}

func (r *OrganizationRepo) Update(ctx context.Context, id string, name *string, subscription *organizationv1.Subscription) (*organizationv1.Organization, error) {
	model := new(models.Organization)

	if _, err := get(ctx, r.db.NewSelect().Model(model), model, id); err != nil {
		return nil, err
	}

	if name != nil {
		model.Name = *name
	}

	if subscription != nil {
		model.Subscription = int32(*subscription)
	}

	return update(ctx, r.db, model, "name", "subscription")
}

func (r *OrganizationRepo) Delete(ctx context.Context, id string) error {
	return delete(ctx, r.db, (*models.Organization)(nil), id)
}
