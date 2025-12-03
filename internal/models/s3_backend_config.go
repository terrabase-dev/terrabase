package models

import (
	"errors"
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

func S3BackendConfigFromProto(s3BackendConfig *s3backendconfigv1.S3BackendConfig) (*S3BackendConfig, error) {
	if s3BackendConfig == nil {
		return &S3BackendConfig{}, nil
	}

	res := &S3BackendConfig{
		ID:          s3BackendConfig.GetId(),
		WorkspaceID: s3BackendConfig.GetWorkspaceId(),
		Bucket:      s3BackendConfig.GetBucket(),
		Key:         s3BackendConfig.GetKey(),
		Region:      s3BackendConfig.GetRegion(),
		Encrypt:     s3BackendConfig.GetEncrypt(),
	}

	if s3BackendConfig.DynamodbLock && s3BackendConfig.S3Lock {
		return nil, errors.New("cannot provide both dynamodb_lock and s3_lock")
	}

	if s3BackendConfig.DynamodbLock {
		if s3BackendConfig.DynamodbTable == nil {
			return nil, errors.New("must provide dynamodb_table if dynamodb_lock is true")
		}

		res.DynamoDBLock = s3BackendConfig.GetDynamodbLock()
		res.DynamoDBTable = s3BackendConfig.GetDynamodbTable()
	} else if s3BackendConfig.S3Lock {
		res.S3Lock = s3BackendConfig.GetS3Lock()
	}

	return res, nil
}

func (s *S3BackendConfig) ToProto() (*s3backendconfigv1.S3BackendConfig, error) {
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

	if s.DynamoDBLock && s.S3Lock {
		return nil, errors.New("cannot provide both dynamodb_lock and s3_lock")
	}

	if s.DynamoDBLock {
		if s.DynamoDBTable == "" {
			return nil, errors.New("must provide dynamodb_table if dynamodb_lock is true")
		}

		res.DynamodbLock = s.DynamoDBLock
		res.DynamodbTable = &s.DynamoDBTable
	} else if s.S3Lock {
		res.S3Lock = s.S3Lock
	}

	return res, nil
}
