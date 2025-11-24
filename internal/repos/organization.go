package repos

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
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
		model.ID = uuid.NewString()
	}

	_, err := r.db.NewInsert().Model(model).Returning("*").Exec(ctx)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("organization: %w", ErrAlreadyExists)
		}

		return nil, err
	}

	return model.ToProto(), nil
}

func (r *OrganizationRepo) Get(ctx context.Context, id string) (*organizationv1.Organization, error) {
	var model models.Organization

	err := r.db.NewSelect().Model(&model).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, fmt.Errorf("organization: %w", ErrNotFound)
		}

		return nil, err
	}

	return model.ToProto(), nil
}

func (r *OrganizationRepo) List(ctx context.Context, pageSize int32, pageToken string) ([]*organizationv1.Organization, string, error) {
	limit := int(pageSize)

	if limit <= 0 || limit > 1000 {
		limit = 50
	}

	offset := 0

	if pageToken != "" {
		if parsed, err := strconv.Atoi(pageToken); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	var models []models.Organization

	err := r.db.NewSelect().Model(&models).Order("created_at DESC").Limit(limit).Offset(offset).Scan(ctx)
	if err != nil {
		return nil, "", err
	}

	orgs := make([]*organizationv1.Organization, 0, len(models))

	for i := range models {
		orgs = append(orgs, models[i].ToProto())
	}

	nextToken := ""

	if len(models) == limit {
		nextToken = strconv.Itoa(offset + limit)
	}

	return orgs, nextToken, nil
}

func (r *OrganizationRepo) Update(ctx context.Context, id string, name *string, subscription *organizationv1.Subscription) (*organizationv1.Organization, error) {
	var model models.Organization

	err := r.db.NewSelect().Model(&model).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, fmt.Errorf("organization: %w", ErrNotFound)
		}

		return nil, err
	}

	if name != nil {
		model.Name = *name
	}

	if subscription != nil {
		model.Subscription = int32(*subscription)
	}

	model.UpdatedAt = time.Now().UTC()

	_, err = r.db.NewUpdate().Model(&model).Column("name", "subscription", "updated_at").WherePK().Exec(ctx)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("organization: %w", ErrAlreadyExists)
		}

		return nil, err
	}

	return model.ToProto(), nil
}

func (r *OrganizationRepo) Delete(ctx context.Context, id string) error {
	res, err := r.db.NewDelete().Model((*models.Organization)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	if rowCount(res) == 0 {
		return ErrNotFound
	}

	return nil
}
