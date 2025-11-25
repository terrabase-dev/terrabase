package services

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
	authv1 "github.com/terrabase-dev/terrabase/specs/terrabase/auth/v1"
	organizationv1 "github.com/terrabase-dev/terrabase/specs/terrabase/organization/v1"
)

type OrganizationService struct {
	AuthAware
	repo   *repos.OrganizationRepo
	logger *log.Logger
}

func NewOrganizationService(repo *repos.OrganizationRepo, logger *log.Logger) *OrganizationService {
	return &OrganizationService{
		repo:   repo,
		logger: logger,
	}
}

func (s *OrganizationService) CreateOrganization(ctx context.Context, req *connect.Request[organizationv1.CreateOrganizationRequest]) (*connect.Response[organizationv1.CreateOrganizationResponse], error) {
	if req.Msg.GetName() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("name is required"))
	}

	if req.Msg.GetSubscription() == organizationv1.Subscription_SUBSCRIPTION_UNSPECIFIED {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("subscription is required"))
	}

	if _, err := s.requireAnyScope(ctx, authv1.Scope_SCOPE_ADMIN, authv1.Scope_SCOPE_ORG_WRITE); err != nil {
		return nil, err
	}

	org := &organizationv1.Organization{
		Name:         req.Msg.GetName(),
		Subscription: req.Msg.GetSubscription(),
	}

	created, err := s.repo.Create(ctx, org)
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&organizationv1.CreateOrganizationResponse{
		Organization: created,
	}), nil
}

func (s *OrganizationService) GetOrganization(ctx context.Context, req *connect.Request[organizationv1.GetOrganizationRequest]) (*connect.Response[organizationv1.GetOrganizationResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("id is required"))
	}

	if _, err := s.requireAnyScope(ctx, authv1.Scope_SCOPE_ADMIN, authv1.Scope_SCOPE_ORG_WRITE, authv1.Scope_SCOPE_ORG_READ); err != nil {
		return nil, err
	}

	org, err := s.repo.Get(ctx, req.Msg.GetId())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&organizationv1.GetOrganizationResponse{Organization: org}), nil
}

func (s *OrganizationService) ListOrganizations(ctx context.Context, req *connect.Request[organizationv1.ListOrganizationsRequest]) (*connect.Response[organizationv1.ListOrganizationsResponse], error) {
	if _, err := s.requireAnyScope(ctx, authv1.Scope_SCOPE_ADMIN, authv1.Scope_SCOPE_ORG_WRITE, authv1.Scope_SCOPE_ORG_READ); err != nil {
		return nil, err
	}

	orgs, nextToken, err := s.repo.List(ctx, req.Msg.GetPageSize(), req.Msg.GetPageToken())
	if err != nil {
		return nil, mapError(err)
	}

	var next *string

	if nextToken != "" {
		next = &nextToken
	}

	return connect.NewResponse(&organizationv1.ListOrganizationsResponse{
		Organizations: orgs,
		NextPageToken: next,
	}), nil
}

func (s *OrganizationService) UpdateOrganization(ctx context.Context, req *connect.Request[organizationv1.UpdateOrganizationRequest]) (*connect.Response[organizationv1.UpdateOrganizationResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("id is required"))
	}

	if req.Msg.Name == nil && req.Msg.Subscription == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("no updates provided"))
	}

	if _, err := s.requireAnyScope(ctx, authv1.Scope_SCOPE_ADMIN, authv1.Scope_SCOPE_ORG_WRITE); err != nil {
		return nil, err
	}

	updated, err := s.repo.Update(ctx, req.Msg.GetId(), req.Msg.Name, req.Msg.Subscription)

	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&organizationv1.UpdateOrganizationResponse{
		Organization: updated,
	}), nil
}

func (s *OrganizationService) DeleteOrganization(ctx context.Context, req *connect.Request[organizationv1.DeleteOrganizationRequest]) (*connect.Response[organizationv1.DeleteOrganizationResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("id is required"))
	}

	if _, err := s.requireAnyScope(ctx, authv1.Scope_SCOPE_ADMIN, authv1.Scope_SCOPE_ORG_WRITE); err != nil {
		return nil, err
	}

	if err := s.repo.Delete(ctx, req.Msg.GetId()); err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&organizationv1.DeleteOrganizationResponse{}), nil
}
