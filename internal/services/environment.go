package services

import (
	"context"
	"log"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
	environmentv1 "github.com/terrabase-dev/terrabase/specs/terrabase/environment/v1"
	workspacev1 "github.com/terrabase-dev/terrabase/specs/terrabase/workspace/v1"
	"github.com/uptrace/bun"
)

type EnvironmentWorkspaceCreator interface {
	CreateForEnvironment(
		ctx context.Context,
		tx bun.Tx,
	)
}

type EnvironmentService struct {
	db               *bun.DB
	repo             *repos.EnvironmentRepo
	workspaceCreator IWorkspaceCreator
	logger           *log.Logger
}

func NewEnvironmentService(
	db *bun.DB,
	repo *repos.EnvironmentRepo,
	workspaceRepo *repos.WorkspaceRepo,
	accessGrantRepo *repos.TeamWorkspaceAccessGrantRepo,
	s3BackendConfigRepo *repos.S3BackendConfigRepo,
	logger *log.Logger,
) *EnvironmentService {
	return &EnvironmentService{
		db:               db,
		repo:             repo,
		workspaceCreator: NewWorkspaceCreator(db, workspaceRepo, accessGrantRepo, s3BackendConfigRepo),
		logger:           logger,
	}
}

func (s *EnvironmentService) CreateEnvironment(ctx context.Context, req *connect.Request[environmentv1.CreateEnvironmentRequest]) (*connect.Response[environmentv1.CreateEnvironmentResponse], error) {
	if req.Msg.GetName() == "" {
		return nil, ErrNameRequired
	}

	if req.Msg.GetApplicationId() == "" {
		return nil, fieldRequiredError("application_id")
	}

	if req.Msg.GetNewWorkspace() == nil {
		return nil, fieldRequiredError("new_workspace")
	}

	environment := &environmentv1.Environment{
		Name:          req.Msg.GetName(),
		ApplicationId: req.Msg.GetApplicationId(),
	}

	var createdEnvironment *environmentv1.Environment

	txErr := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		environmentRepo := s.repo.WithTx(tx)
		var err error

		createdEnvironment, err = environmentRepo.Create(ctx, environment)
		if err != nil {
			return err
		}

		createWorkspaceReq := req.Msg.GetNewWorkspace()

		createWorkspaceReq.Owner = &workspacev1.CreateWorkspaceRequest_EnvironmentId{EnvironmentId: createdEnvironment.Id}

		if _, err := s.workspaceCreator.CreateWorkspace(ctx, tx, createWorkspaceReq); err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, mapError(txErr)
	}

	return connect.NewResponse(&environmentv1.CreateEnvironmentResponse{Environment: createdEnvironment}), nil
}

func (s *EnvironmentService) GetEnvironment(ctx context.Context, req *connect.Request[environmentv1.GetEnvironmentRequest]) (*connect.Response[environmentv1.GetEnvironmentResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, ErrIdRequired
	}

	environment, err := s.repo.Get(ctx, req.Msg.GetId())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&environmentv1.GetEnvironmentResponse{Environment: environment}), nil
}

func (s *EnvironmentService) ListEnvironments(ctx context.Context, req *connect.Request[environmentv1.ListEnvironmentsRequest]) (*connect.Response[environmentv1.ListEnvironmentsResponse], error) {
	if req.Msg.GetApplicationId() == "" {
		return nil, fieldRequiredError("application_id")
	}

	environments, err := s.repo.List(ctx, req.Msg.GetApplicationId())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&environmentv1.ListEnvironmentsResponse{Environments: environments}), nil
}

func (s *EnvironmentService) UpdateEnvironment(ctx context.Context, req *connect.Request[environmentv1.UpdateEnvironmentRequest]) (*connect.Response[environmentv1.UpdateEnvironmentResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, ErrIdRequired
	}

	if req.Msg.GetName() == "" {
		return nil, ErrNameRequired
	}

	updated, err := s.repo.Update(ctx, req.Msg.GetId(), req.Msg.GetName())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&environmentv1.UpdateEnvironmentResponse{Environment: updated}), nil
}

func (s *EnvironmentService) DeleteEnvironment(ctx context.Context, req *connect.Request[environmentv1.DeleteEnvironmentRequest]) (*connect.Response[environmentv1.DeleteEnvironmentResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, ErrIdRequired
	}

	if err := s.repo.Delete(ctx, req.Msg.GetId()); err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&environmentv1.DeleteEnvironmentResponse{}), nil
}
