package services

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
	applicationv1 "github.com/terrabase-dev/terrabase/specs/terrabase/application/v1"
	teamAccessTypev1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_access_type/v1"
	teamApplicationAccessGrantv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team_application_access_grant/v1"
)

type ApplicationService struct {
	AuthAware
	repo            *repos.ApplicationRepo
	accessGrantRepo *repos.TeamApplicationAccessGrantRepo
	logger          *log.Logger
}

func NewApplicationService(repo *repos.ApplicationRepo, logger *log.Logger) *ApplicationService {
	return &ApplicationService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ApplicationService) CreateApplication(ctx context.Context, req *connect.Request[applicationv1.CreateApplicationRequest]) (*connect.Response[applicationv1.CreateApplicationResponse], error) {
	if req.Msg.GetName() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("name is required"))
	}

	app := &applicationv1.Application{
		Name: req.Msg.GetName(),
	}

	application, err := s.repo.Create(ctx, app)
	if err != nil {
		return nil, mapError(err)
	}

	teamApplicationAccessGrant := &teamApplicationAccessGrantv1.TeamApplicationAccessGrant{
		ApplicationId: application.Id,
		TeamId:        req.Msg.GetTeamId(),
		AccessType:    teamAccessTypev1.TeamAccessType_TEAM_ACCESS_TYPE_OWNER,
	}

	if _, err := s.accessGrantRepo.Create(ctx, teamApplicationAccessGrant); err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&applicationv1.CreateApplicationResponse{
		Application: application,
	}), nil
}

func (s *ApplicationService) GetApplication(ctx context.Context, req *connect.Request[applicationv1.GetApplicationRequest]) (*connect.Response[applicationv1.GetApplicationResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, IDRequiredError
	}

	app, err := s.repo.Get(ctx, req.Msg.GetId())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&applicationv1.GetApplicationResponse{Application: app}), nil
}

func (s *ApplicationService) ListApplications(ctx context.Context, req *connect.Request[applicationv1.ListApplicationsRequest]) (*connect.Response[applicationv1.ListApplicationsResponse], error) {
	if req.Msg.GetTeamId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("team_id is required"))
	}

	applications, nextToken, err := s.repo.List(ctx, req.Msg.GetTeamId(), req.Msg.GetPageSize(), req.Msg.GetPageToken())
	if err != nil {
		return nil, mapError(err)
	}

	var next *string

	if nextToken != "" {
		next = &nextToken
	}

	return connect.NewResponse(&applicationv1.ListApplicationsResponse{
		Applications:  applications,
		NextPageToken: next,
	}), nil
}

func (s *ApplicationService) UpdateApplication(ctx context.Context, req *connect.Request[applicationv1.UpdateApplicationRequest]) (*connect.Response[applicationv1.UpdateApplicationResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, IDRequiredError
	}

	if req.Msg.Name == "" {
		return nil, NoUpdatesProvidedError
	}

	updated, err := s.repo.Update(ctx, req.Msg.GetId(), req.Msg.GetName())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&applicationv1.UpdateApplicationResponse{Application: updated}), nil
}

func (s *ApplicationService) DeleteApplication(ctx context.Context, req *connect.Request[applicationv1.DeleteApplicationRequest]) (*connect.Response[applicationv1.DeleteApplicationResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, IDRequiredError
	}

	if err := s.repo.Delete(ctx, req.Msg.GetId()); err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&applicationv1.DeleteApplicationResponse{}), nil
}

func (s *ApplicationService) GrantTeamAccess(ctx context.Context, req *connect.Request[applicationv1.GrantTeamAccessRequest]) (*connect.Response[applicationv1.GrantTeamAccessResponse], error) {
	teamAccessGrants := req.Msg.GetTeamAccessGrants()

	for i := range teamAccessGrants {
		teamAccessGrant := &teamApplicationAccessGrantv1.TeamApplicationAccessGrant{
			TeamId:        teamAccessGrants[i].GetTeamId(),
			ApplicationId: teamAccessGrants[i].GetApplicationId(),
			AccessType:    teamAccessGrants[i].GetAccessType(),
		}

		if _, err := s.accessGrantRepo.Create(ctx, teamAccessGrant); err != nil {
			return nil, mapError(err)
		}
	}

	return connect.NewResponse(&applicationv1.GrantTeamAccessResponse{}), nil
}

func (s *ApplicationService) RevokeTeamAccess(ctx context.Context, req *connect.Request[applicationv1.RevokeTeamAccessRequest]) (*connect.Response[applicationv1.RevokeTeamAccessResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, IDRequiredError
	}

	applicationId := req.Msg.GetId()
	teamIds := req.Msg.GetTeamIds().GetTeamIds()

	for i := range teamIds {
		if err := s.accessGrantRepo.Delete(ctx, teamIds[i], applicationId); err != nil {
			return nil, mapError(err)
		}
	}

	return connect.NewResponse(&applicationv1.RevokeTeamAccessResponse{}), nil
}
