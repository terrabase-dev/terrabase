package repos

import (
	"context"

	"github.com/terrabase-dev/terrabase/internal/models"
	s3backendconfigv1 "github.com/terrabase-dev/terrabase/specs/terrabase/s3_backend_config/v1"
	"github.com/uptrace/bun"
)

type S3BackendConfigRepo struct {
	db *bun.DB
}

func NewS3BackendConfigRepo(db *bun.DB) *S3BackendConfigRepo {
	return &S3BackendConfigRepo{
		db: db,
	}
}

func (r *S3BackendConfigRepo) Create(ctx context.Context, s3BackendConfig *s3backendconfigv1.S3BackendConfig) (*s3backendconfigv1.S3BackendConfig, error) {
	model := models.S3BackendConfigFromProto(s3BackendConfig)

	if model.ID == "" {
		model.ID = uuidString()
	}

	return create(ctx, r.db, model)
}

func (r *S3BackendConfigRepo) Get(ctx context.Context, id string) (*s3backendconfigv1.S3BackendConfig, error) {
	model := new(models.S3BackendConfig)

	return get(ctx, r.db.NewSelect().Model(model), model, id)
}

func (r *S3BackendConfigRepo) Update(ctx context.Context, id string, workspaceId *string, bucket *string, key *string, region *string, dynamodbLock *bool, s3Lock *bool, encrypt *bool, dynamodbTable *string) (*s3backendconfigv1.S3BackendConfig, error) {
	model := new(models.S3BackendConfig)

	if _, err := get(ctx, r.db.NewSelect().Model(model), model, id); err != nil {
		return nil, err
	}

	model.WorkspaceID = *workspaceId
	model.Bucket = *bucket
	model.Key = *key
	model.Region = *region
	model.DynamoDBLock = *dynamodbLock
	model.S3Lock = *s3Lock
	model.Encrypt = *encrypt
	model.DynamoDBTable = *dynamodbTable

	return update(ctx, r.db, model, "workspace_id", "bucket", "key", "region", "dynamodb_lock", "s3_lock", "encrypt", "dynamodb_table")
}

func (r *S3BackendConfigRepo) Delete(ctx context.Context, id string) error {
	return delete(ctx, r.db, (*models.S3BackendConfig)(nil), id)
}
