package services

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
	teamv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team/v1"
)

type TeamService struct {
	AuthAware
	repo   *repos.TeamRepo
	logger *log.Logger
}

func NewTeamService(repo *repos.TeamRepo, logger *log.Logger) *TeamService {
	return &TeamService{
		repo:   repo,
		logger: logger,
	}
}

func (s *TeamService) CreateTeam(ctx context.Context, req *connect.Request[teamv1.CreateTeamRequest]) (*connect.Response[teamv1.CreateTeamResponse], error) {
	if req.Msg.GetName() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("name is required"))
	}

	if req.Msg.OrganizationId == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("organization_id is required"))
	}

	team := &teamv1.Team{
		Name:           req.Msg.GetName(),
		OrganizationId: req.Msg.GetOrganizationId(),
	}

	created, err := s.repo.Create(ctx, team)
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&teamv1.CreateTeamResponse{
		Team: created,
	}), nil
}

func (s *TeamService) GetTeam(ctx context.Context, req *connect.Request[teamv1.GetTeamRequest]) (*connect.Response[teamv1.GetTeamResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("id is required"))
	}

	team, err := s.repo.Get(ctx, req.Msg.GetId())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&teamv1.GetTeamResponse{Team: team}), nil
}

func (s *TeamService) ListTeams(ctx context.Context, req *connect.Request[teamv1.ListTeamsRequest]) (*connect.Response[teamv1.ListTeamsResponse], error) {
	teams, nextToken, err := s.repo.List(ctx, req.Msg.GetPageSize(), req.Msg.GetPageToken())
	if err != nil {
		return nil, mapError(err)
	}

	var next *string

	if nextToken != "" {
		next = &nextToken
	}

	return connect.NewResponse(&teamv1.ListTeamsResponse{
		Teams:         teams,
		NextPageToken: next,
	}), nil
}

func (s *TeamService) UpdateTeam(ctx context.Context, req *connect.Request[teamv1.UpdateTeamRequest]) (*connect.Response[teamv1.UpdateTeamResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("id is required"))
	}

	if req.Msg.GetName() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("no updates provided"))
	}

	updated, err := s.repo.Update(ctx, req.Msg.GetId(), req.Msg.GetName())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&teamv1.UpdateTeamResponse{
		Team: updated,
	}), nil
}

func (s *TeamService) DeleteTeam(ctx context.Context, req *connect.Request[teamv1.DeleteTeamRequest]) (*connect.Response[teamv1.DeleteTeamResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("id is required"))
	}

	if err := s.repo.Delete(ctx, req.Msg.GetId()); err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&teamv1.DeleteTeamResponse{}), nil
}
