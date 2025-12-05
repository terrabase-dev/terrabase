package repos

import (
	"context"
	"errors"

	"github.com/terrabase-dev/terrabase/internal/models"
	s3BackendConfigv1 "github.com/terrabase-dev/terrabase/specs/terrabase/s3_backend_config/v1"
	"github.com/uptrace/bun"
)

type S3BackendConfigRepo struct {
	db bun.IDB
}

func NewS3BackendConfigRepo(db *bun.DB) *S3BackendConfigRepo {
	return &S3BackendConfigRepo{db: db}
}

func (r *S3BackendConfigRepo) WithTx(tx bun.Tx) *S3BackendConfigRepo {
	return &S3BackendConfigRepo{db: tx}
}

func (r *S3BackendConfigRepo) Create(ctx context.Context, s3BackendConfig *s3BackendConfigv1.S3BackendConfig) (*s3BackendConfigv1.S3BackendConfig, error) {
	model := r.createWithId(s3BackendConfig)

	return create(ctx, r.db, model)
}

func (r *S3BackendConfigRepo) CreateForWorkspace(ctx context.Context, s3BackendConfig *s3BackendConfigv1.S3BackendConfig, workspaceId string) (*s3BackendConfigv1.S3BackendConfig, error) {
	model := r.createWithId(s3BackendConfig)

	model.WorkspaceID = workspaceId

	return create(ctx, r.db, model)
}

func (r *S3BackendConfigRepo) Get(ctx context.Context, id string) (*s3BackendConfigv1.S3BackendConfig, error) {
	model := new(models.S3BackendConfig)

	return get(ctx, r.db.NewSelect().Model(model), model, id)
}

func (r *S3BackendConfigRepo) Update(ctx context.Context, id string, workspaceId *string, bucket *string, key *string, region *string, dynamodbLock *bool, s3Lock *bool, encrypt *bool, dynamodbTable *string) (*s3BackendConfigv1.S3BackendConfig, error) {
	model := new(models.S3BackendConfig)

	if _, err := get(ctx, r.db.NewSelect().Model(model), model, id); err != nil {
		return nil, err
	}

	colsToUpdate := make([]string, 0, 8)

	if workspaceId != nil {
		model.WorkspaceID = *workspaceId
		colsToUpdate = append(colsToUpdate, "workspace_id")
	}

	if bucket != nil {
		model.Bucket = *bucket
		colsToUpdate = append(colsToUpdate, "bucket")
	}

	if key != nil {
		model.Key = *key
		colsToUpdate = append(colsToUpdate, "key")
	}

	if region != nil {
		model.Region = *region
		colsToUpdate = append(colsToUpdate, "region")
	}

	if dynamodbLock != nil {
		model.DynamoDBLock = *dynamodbLock
		colsToUpdate = append(colsToUpdate, "dynamodb_lock")
	}

	if s3Lock != nil {
		model.S3Lock = *s3Lock
		colsToUpdate = append(colsToUpdate, "s3_lock")
	}

	if encrypt != nil {
		model.Encrypt = *encrypt
		colsToUpdate = append(colsToUpdate, "encrypt")
	}

	if dynamodbTable != nil {
		model.DynamoDBTable = *dynamodbTable
		colsToUpdate = append(colsToUpdate, "dynamodb_table")
	}

	if len(colsToUpdate) == 0 {
		return nil, errors.New("no columns to update")
	}

	return update(ctx, r.db, model, colsToUpdate...)
}

func (r *S3BackendConfigRepo) Delete(ctx context.Context, id string) error {
	return delete(ctx, r.db, (*models.S3BackendConfig)(nil), id)
}

func (r *S3BackendConfigRepo) createWithId(s3BackendConfig *s3BackendConfigv1.S3BackendConfig) *models.S3BackendConfig {
	model := models.S3BackendConfigFromProto(s3BackendConfig)

	if model.ID == "" {
		model.ID = uuidString()
	}

	return model
}
