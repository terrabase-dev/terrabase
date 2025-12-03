package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/terrabase-dev/terrabase/internal/models"
	userv1 "github.com/terrabase-dev/terrabase/specs/terrabase/user/v1"
	"github.com/uptrace/bun"
)

type UserRepo struct {
	db *bun.DB
}

func NewUserRepo(db *bun.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *userv1.User) (*userv1.User, error) {
	model := models.UserFromProto(user)

	if model.ID == "" {
		model.ID = uuidString()
	}

	if model.UserType == 0 {
		model.UserType = userv1.UserType_USER_TYPE_USER
	}

	now := time.Now().UTC()

	model.CreatedAt = now
	model.UpdatedAt = now

	_, err := r.db.NewInsert().Model(model).Returning("*").Exec(ctx)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("user: %w", ErrAlreadyExists)
		}

		return nil, err
	}

	return model.ToProto(), nil
}

func (r *UserRepo) Get(ctx context.Context, id string) (*models.User, error) {
	var model models.User

	err := r.db.NewSelect().Model(&model).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, fmt.Errorf("user: %w", ErrNotFound)
		}

		return nil, err
	}

	return &model, nil
}

// TODO: implement
// func (r *UserRepo) List(ctx context.Context, pageSize int32, pageToken string, organization_id string, team_id string, workspace_id string, user_type userv1.UserType) ([]*userv1.UserSummary, string, error) {
//     limit := int(pageSize)

//     if limit <= 0 || limit > 1000 {
//         limit = 50
//     }

//     offset := 0

//     if pageToken != "" {
//         if parsed, err := strconv.Atoi(pageToken); err == nil && parsed >= 0 {
//             offset = parsed
//         }
//     }

//     var models []models.User

//     err := r.db.NewSelect().Model(&models).Order("name ASC").Limit(limit).Offset(offset).Scan(ctx)
//     if err != nil {
//         return nil, "", err
//     }

//     users := make([]*userv1.User, 0, len(models))

//     for i := range(models) {
//         users = append(users, models[i].ToProto())
//     }

//     nextToken := ""

//     if len(models) == limit {
//         nextToken = strconv.Itoa(offset + limit)
//     }

//     return users, nextToken, nil
// }

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var model models.User

	err := r.db.NewSelect().Model(&model).Where("email = ?", email).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, fmt.Errorf("user: %w", ErrNotFound)
		}

		return nil, err
	}

	return &model, nil
}

func (r *UserRepo) Update(ctx context.Context, id string, name *string, email *string, default_role *int32) (*userv1.User, error) {
	var user models.User

	err := r.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, fmt.Errorf("user: %w", ErrNotFound)
		}

		return nil, err
	}

	if name != nil {
		user.Name = *name
	}

	if email != nil {
		user.Email = *email
	}

	if default_role != nil {
		user.DefaultRole = *default_role
	}

	user.UpdatedAt = time.Now().UTC()

	_, err = r.db.NewUpdate().Model(&user).Column("name", "email", "default_role").WherePK().Exec(ctx)
	if err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("user: %w", ErrAlreadyExists)
		}

		return nil, err
	}

	return user.ToProto(), nil
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	res, err := r.db.NewDelete().Model((*models.User)(nil)).Where("id = ?", id).Exec(ctx)

	if err != nil {
		return err
	}

	if rowCount(res) == 0 {
		return ErrNotFound
	}

	return nil
}
