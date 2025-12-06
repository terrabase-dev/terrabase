package services

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
	teamAccessTypev1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_access_type/v1"
	teamWorkspaceAccessGrantv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_workspace_access_grant/v1"
	workspacev1 "github.com/terrabase-dev/terrabase/specs/terrabase/workspace/v1"
	"github.com/uptrace/bun"
)

type IWorkspaceCreator interface {
	CreateWorkspace(
		ctx context.Context,
		tx bun.Tx,
		msg *workspacev1.CreateWorkspaceRequest,
	) (*workspacev1.Workspace, error)
}

type WorkspaceBackendCreator interface {
	CreateForWorkspace(
		ctx context.Context,
		tx bun.Tx,
		workspace *workspacev1.Workspace,
		msg *workspacev1.CreateWorkspaceRequest,
	) error
}

type WorkspaceService struct {
	AuthAware
	db               *bun.DB
	repo             *repos.WorkspaceRepo
	accessGrantRepo  *repos.TeamWorkspaceAccessGrantRepo
	workspaceCreator IWorkspaceCreator
	logger           *log.Logger
}

func NewWorkspaceService(
	db *bun.DB,
	repo *repos.WorkspaceRepo,
	accessGrantRepo *repos.TeamWorkspaceAccessGrantRepo,
	s3BackendConfigRepo *repos.S3BackendConfigRepo,
	logger *log.Logger,
) *WorkspaceService {
	return &WorkspaceService{
		db:               db,
		repo:             repo,
		accessGrantRepo:  accessGrantRepo,
		workspaceCreator: NewWorkspaceCreator(db, repo, accessGrantRepo, s3BackendConfigRepo),
		logger:           logger,
	}
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, req *connect.Request[workspacev1.CreateWorkspaceRequest]) (*connect.Response[workspacev1.CreateWorkspaceResponse], error) {
	var createdWorkspace *workspacev1.Workspace

	txErr := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		var err error

		createdWorkspace, err = s.workspaceCreator.CreateWorkspace(ctx, tx, req.Msg)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, mapError(txErr)
	}

	return connect.NewResponse(&workspacev1.CreateWorkspaceResponse{Workspace: createdWorkspace}), nil
}

func (s *WorkspaceService) GetWorkspace(ctx context.Context, req *connect.Request[workspacev1.GetWorkspaceRequest]) (*connect.Response[workspacev1.GetWorkspaceResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, ErrIdRequired
	}

	workspace, err := s.repo.Get(ctx, req.Msg.GetId())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&workspacev1.GetWorkspaceResponse{Workspace: workspace}), nil
}

func (s *WorkspaceService) ListWorkspaces(ctx context.Context, req *connect.Request[workspacev1.ListWorkspacesRequest]) (*connect.Response[workspacev1.ListWorkspacesResponse], error) {
	workspaces, nextToken, err := s.repo.List(ctx, req.Msg.TeamId, req.Msg.ApplicationId, req.Msg.GetPageSize(), req.Msg.GetPageToken())
	if err != nil {
		return nil, mapError(err)
	}

	var next *string

	if nextToken != "" {
		next = &nextToken
	}

	return connect.NewResponse(&workspacev1.ListWorkspacesResponse{
		Workspaces:    workspaces,
		NextPageToken: next,
	}), nil
}

func (s *WorkspaceService) UpdateWorkspace(ctx context.Context, req *connect.Request[workspacev1.UpdateWorkspaceRequest]) (*connect.Response[workspacev1.UpdateWorkspaceResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, ErrIdRequired
	}

	if req.Msg.EnvironmentId != nil && req.Msg.TeamId != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("cannot provide both environment_id and team_id"))
	}

	updated, err := s.repo.Update(ctx, req.Msg.GetId(), req.Msg.Name, req.Msg.BackendType, req.Msg.EnvironmentId)
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&workspacev1.UpdateWorkspaceResponse{Workspace: updated}), nil
}

func (s *WorkspaceService) DeleteWorkspace(ctx context.Context, req *connect.Request[workspacev1.DeleteWorkspaceRequest]) (*connect.Response[workspacev1.DeleteWorkspaceResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, ErrIdRequired
	}

	if err := s.repo.Delete(ctx, req.Msg.GetId()); err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&workspacev1.DeleteWorkspaceResponse{}), nil
}

func (s *WorkspaceService) GrantTeamAccess(ctx context.Context, req *connect.Request[workspacev1.GrantTeamAccessRequest]) (*connect.Response[workspacev1.GrantTeamAccessResponse], error) {
	teamAccessGrants := req.Msg.GetTeamAccessGrants()

	for i := range teamAccessGrants {
		teamAccessGrant := &teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant{
			TeamId:      teamAccessGrants[i].GetTeamId(),
			WorkspaceId: teamAccessGrants[i].GetWorkspaceId(),
			AccessType:  teamAccessGrants[i].GetAccessType(),
		}

		if _, err := s.accessGrantRepo.Create(ctx, teamAccessGrant); err != nil {
			return nil, mapError(err)
		}
	}

	return connect.NewResponse(&workspacev1.GrantTeamAccessResponse{}), nil
}

func (s *WorkspaceService) RevokeTeamAccess(ctx context.Context, req *connect.Request[workspacev1.RevokeTeamAccessRequest]) (*connect.Response[workspacev1.RevokeTeamAccessResponse], error) {
	if req.Msg.GetWorkspaceId() == "" {
		return nil, fieldRequiredError("workspace_id")
	}

	workspaceId := req.Msg.GetWorkspaceId()
	teamIds := req.Msg.GetTeamIds().GetTeamIds()

	for i := range teamIds {
		if err := s.accessGrantRepo.Delete(ctx, teamIds[i], workspaceId); err != nil {
			return nil, mapError(err)
		}
	}

	return connect.NewResponse(&workspacev1.RevokeTeamAccessResponse{}), nil
}

func buildWorkspaceFromCreateReq(msg *workspacev1.CreateWorkspaceRequest) (*workspacev1.Workspace, string, error) {
	if msg.GetName() == "" {
		return nil, "", ErrNameRequired
	}

	if msg.GetBackendType() == workspacev1.BackendType_BACKEND_TYPE_UNSPECIFIED {
		return nil, "", fieldRequiredError("backend_type")
	}

	if msg.GetBackendType() == workspacev1.BackendType_BACKEND_TYPE_S3 && msg.GetS3BackendConfig() == nil {
		return nil, "", connect.NewError(connect.CodeInvalidArgument, errors.New("must provide s3_backend_config when backend_type is S3"))
	}

	var environmentId, teamId string

	switch l := msg.Owner.(type) {
	case *workspacev1.CreateWorkspaceRequest_EnvironmentId:
		environmentId = l.EnvironmentId
		teamId = ""
	case *workspacev1.CreateWorkspaceRequest_TeamId:
		teamId = l.TeamId
		environmentId = ""
	case nil:
		return nil, "", connect.NewError(connect.CodeInvalidArgument, errors.New("must provide exactly one of environment_id or team_id"))
	default:
		return nil, "", connect.NewError(connect.CodeInvalidArgument, errors.New("must provide exactly one of environment_id or team_id"))
	}

	workspace := &workspacev1.Workspace{
		Name:        msg.GetName(),
		BackendType: msg.GetBackendType(),
	}

	if environmentId != "" {
		workspace.EnvironmentId = &environmentId
	}

	return workspace, teamId, nil
}

type WorkspaceCreator struct {
	db              *bun.DB
	repo            *repos.WorkspaceRepo
	accessGrantRepo *repos.TeamWorkspaceAccessGrantRepo
	backendCreators map[workspacev1.BackendType]WorkspaceBackendCreator
}

func NewWorkspaceCreator(db *bun.DB, repo *repos.WorkspaceRepo, accessGrantRepo *repos.TeamWorkspaceAccessGrantRepo, s3BackendConfigRepo *repos.S3BackendConfigRepo) *WorkspaceCreator {
	return &WorkspaceCreator{
		db:              db,
		repo:            repo,
		accessGrantRepo: accessGrantRepo,
		backendCreators: map[workspacev1.BackendType]WorkspaceBackendCreator{
			workspacev1.BackendType_BACKEND_TYPE_S3: NewS3BackendCreator(s3BackendConfigRepo),
		},
	}
}

func (c *WorkspaceCreator) CreateWorkspace(ctx context.Context, tx bun.Tx, msg *workspacev1.CreateWorkspaceRequest) (*workspacev1.Workspace, error) {
	workspace, teamId, err := buildWorkspaceFromCreateReq(msg)
	if err != nil {
		return nil, err
	}

	var createdWorkspace *workspacev1.Workspace
	backendCreator, ok := c.backendCreators[msg.GetBackendType()]
	if !ok {
		return nil, errors.New("unsupported backend_type")
	}

	workspaceRepo := c.repo.WithTx(tx)

	createdWorkspace, err = workspaceRepo.Create(ctx, workspace)
	if err != nil {
		return nil, err
	}

	if err := backendCreator.CreateForWorkspace(ctx, tx, createdWorkspace, msg); err != nil {
		return nil, err
	}

	if teamId != "" {
		accessGrantRepo := c.accessGrantRepo.WithTx(tx)

		teamWorkspaceAccessGrant := &teamWorkspaceAccessGrantv1.TeamWorkspaceAccessGrant{
			TeamId:      teamId,
			WorkspaceId: createdWorkspace.Id,
			AccessType:  teamAccessTypev1.TeamAccessType_TEAM_ACCESS_TYPE_OWNER,
		}

		if _, err := accessGrantRepo.Create(ctx, teamWorkspaceAccessGrant); err != nil {
			return nil, err
		}
	}

	return createdWorkspace, nil
}
