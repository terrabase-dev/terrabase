package models

import (
	"time"

	s3backendconfigv1 "github.com/terrabase-dev/terrabase/specs/terrabase/s3_backend_config/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type S3BackendConfig struct {
	bun.BaseModel `bun:"table:s3_backend_config"`

	ID            string    `bun:",pk"`
	WorkspaceID   string    `bun:",unique,notnull,on_delete:CASCADE"`
	Bucket        string    `bun:",notnull"`
	Key           string    `bun:",notnull"`
	Region        string    `bun:",notnull"`
	DynamoDBLock  bool      `bun:",default:false"`
	S3Lock        bool      `bun:",default:true"`
	Encrypt       bool      `bun:",default:true"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DynamoDBTable string

	Workspace *Workspace `bun:"rel:belongs-to,join:workspace_id=id"`
}

func S3BackendConfigFromProto(s3BackendConfig *s3backendconfigv1.S3BackendConfig) *S3BackendConfig {
	if s3BackendConfig == nil {
		return &S3BackendConfig{}
	}

	res := &S3BackendConfig{
		ID:          s3BackendConfig.GetId(),
		WorkspaceID: s3BackendConfig.GetWorkspaceId(),
		Bucket:      s3BackendConfig.GetBucket(),
		Key:         s3BackendConfig.GetKey(),
		Region:      s3BackendConfig.GetRegion(),
		Encrypt:     s3BackendConfig.GetEncrypt(),
	}

	if s3BackendConfig.DynamodbLock {
		res.DynamoDBLock = true
		res.DynamoDBTable = s3BackendConfig.GetDynamodbTable()
		res.S3Lock = false
	} else if s3BackendConfig.S3Lock {
		res.S3Lock = true
		res.DynamoDBLock = false
	}

	return res
}

func (s *S3BackendConfig) ToProto() *s3backendconfigv1.S3BackendConfig {
	res := &s3backendconfigv1.S3BackendConfig{
		Id:          s.ID,
		WorkspaceId: s.WorkspaceID,
		Bucket:      s.Bucket,
		Key:         s.Key,
		Region:      s.Region,
		Encrypt:     s.Encrypt,
		CreatedAt:   timestamppb.New(s.CreatedAt.UTC()),
		UpdatedAt:   timestamppb.New(s.UpdatedAt.UTC()),
	}

	if s.DynamoDBLock {
		res.DynamodbLock = true
		res.DynamodbTable = &s.DynamoDBTable
		res.S3Lock = false
	} else if s.S3Lock {
		res.S3Lock = true
		res.DynamodbLock = false
	}

	return res
}

func (a *S3BackendConfig) SetUpdatedAt(updatedAt time.Time) {
	a.UpdatedAt = updatedAt
}

func (a *S3BackendConfig) ModelName() string {
	return "s3_backend_config"
}
