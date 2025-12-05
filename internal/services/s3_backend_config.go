package services

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
	s3BackendConfigv1 "github.com/terrabase-dev/terrabase/specs/terrabase/s3_backend_config/v1"
	workspacev1 "github.com/terrabase-dev/terrabase/specs/terrabase/workspace/v1"
	"github.com/uptrace/bun"
)

type S3BackendConfigService struct {
	AuthAware
	repo   *repos.S3BackendConfigRepo
	logger *log.Logger
}

func NewS3BackendConfigService(repo *repos.S3BackendConfigRepo, logger *log.Logger) *S3BackendConfigService {
	return &S3BackendConfigService{
		repo:   repo,
		logger: logger,
	}
}

func (s *S3BackendConfigService) CreateS3BackendConfig(ctx context.Context, req *connect.Request[s3BackendConfigv1.CreateS3BackendConfigRequest]) (*connect.Response[s3BackendConfigv1.CreateS3BackendConfigResponse], error) {
	s3BackendConfig, err := buildS3BackendConfigFromCreateReq(req.Msg)
	if err != nil {
		return nil, err
	}

	created, err := s.repo.Create(ctx, s3BackendConfig)
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&s3BackendConfigv1.CreateS3BackendConfigResponse{S3BackendConfig: created}), nil
}

func (s *S3BackendConfigService) GetS3BackendConfig(ctx context.Context, req *connect.Request[s3BackendConfigv1.GetS3BackendConfigRequest]) (*connect.Response[s3BackendConfigv1.GetS3BackendConfigResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, ErrIdRequired
	}

	s3BackendConfig, err := s.repo.Get(ctx, req.Msg.GetId())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&s3BackendConfigv1.GetS3BackendConfigResponse{S3BackendConfig: s3BackendConfig}), nil
}

func (s *S3BackendConfigService) UpdateS3BackendConfig(ctx context.Context, req *connect.Request[s3BackendConfigv1.UpdateS3BackendConfigRequest]) (*connect.Response[s3BackendConfigv1.UpdateS3BackendConfigResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, ErrIdRequired
	}

	if req.Msg.DynamodbLock != nil && req.Msg.S3Lock != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("cannot provide both dynamodb_lock and s3_lock"))
	}

	if req.Msg.DynamodbLock != nil && req.Msg.GetDynamodbLock() && req.Msg.GetDynamodbTable() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("must provide dynamodb_table when dynamodb_lock is true"))
	}

	updated, err := s.repo.Update(ctx, req.Msg.GetId(), req.Msg.WorkspaceId, req.Msg.Bucket, req.Msg.Key, req.Msg.Region, req.Msg.DynamodbLock, req.Msg.S3Lock, req.Msg.Encrypt, req.Msg.DynamodbTable)
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&s3BackendConfigv1.UpdateS3BackendConfigResponse{S3BackendConfig: updated}), nil
}

func (s *S3BackendConfigService) DeleteS3BackendConfig(ctx context.Context, req *connect.Request[s3BackendConfigv1.DeleteS3BackendConfigRequest]) (*connect.Response[s3BackendConfigv1.DeleteS3BackendConfigResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, ErrIdRequired
	}

	if err := s.repo.Delete(ctx, req.Msg.GetId()); err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&s3BackendConfigv1.DeleteS3BackendConfigResponse{}), nil
}

func buildS3BackendConfigFromCreateReq(msg *s3BackendConfigv1.CreateS3BackendConfigRequest) (*s3BackendConfigv1.S3BackendConfig, error) {
	if msg.GetBucket() == "" {
		return nil, fieldRequiredError("bucket")
	}

	if msg.GetKey() == "" {
		return nil, fieldRequiredError("key")
	}

	if msg.GetRegion() == "" {
		return nil, fieldRequiredError("region")
	}

	var dynamoDbLock, s3Lock bool

	switch l := msg.Lock.(type) {
	case *s3BackendConfigv1.CreateS3BackendConfigRequest_DynamodbLock:
		dynamoDbLock = l.DynamodbLock
		s3Lock = false
	case *s3BackendConfigv1.CreateS3BackendConfigRequest_S3Lock:
		s3Lock = l.S3Lock
		dynamoDbLock = false
	case nil:
		dynamoDbLock = false
		s3Lock = true
	default:
		dynamoDbLock = false
		s3Lock = true
	}

	if dynamoDbLock && msg.GetDynamodbTable() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("must provide dynamodb_table when dynamodb_lock is true"))
	}

	s3BackendConfig := &s3BackendConfigv1.S3BackendConfig{
		Bucket:       msg.GetBucket(),
		Key:          msg.GetKey(),
		Region:       msg.GetRegion(),
		DynamodbLock: dynamoDbLock,
		S3Lock:       s3Lock,
		Encrypt:      msg.GetEncrypt(),
	}

	if msg.GetDynamodbTable() != "" {
		s3BackendConfig.DynamodbTable = msg.DynamodbTable
	}

	return s3BackendConfig, nil
}

type S3BackendCreator struct {
	repo *repos.S3BackendConfigRepo
}

func NewS3BackendCreator(repo *repos.S3BackendConfigRepo) *S3BackendCreator {
	return &S3BackendCreator{repo: repo}
}

func (c *S3BackendCreator) CreateForWorkspace(ctx context.Context, tx bun.Tx, workspace *workspacev1.Workspace, msg *workspacev1.CreateWorkspaceRequest) error {
	s3BackendConfigRepo := c.repo.WithTx(tx)

	s3BackendConfig, err := buildS3BackendConfigFromCreateReq(msg.GetS3BackendConfig())
	if err != nil {
		return err
	}

	if _, err := s3BackendConfigRepo.CreateForWorkspace(ctx, s3BackendConfig, workspace.Id); err != nil {
		return err
	}

	return nil
}
