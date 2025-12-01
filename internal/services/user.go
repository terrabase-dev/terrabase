package services

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
	userv1 "github.com/terrabase-dev/terrabase/specs/terrabase/user/v1"
)

type UserService struct {
	AuthAware
	repo   *repos.UserRepo
	logger *log.Logger
}

func NewUserService(repo *repos.UserRepo, logger *log.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *connect.Request[userv1.GetUserRequest]) (*connect.Response[userv1.GetUserResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("id is required"))
	}

	user, err := s.repo.Get(ctx, req.Msg.GetId())
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&userv1.GetUserResponse{User: user.ToProto()}), nil
}

func (s *UserService) ListUsers(ctx context.Context, req *connect.Request[userv1.ListUsersRequest]) (*connect.Response[userv1.ListUsersResponse], error) {
	return connect.NewResponse(&userv1.ListUsersResponse{}), nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *connect.Request[userv1.UpdateUserRequest]) (*connect.Response[userv1.UpdateUserResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("id is required"))
	}

	updated, err := s.repo.Update(ctx, req.Msg.GetId(), req.Msg.Name, req.Msg.Email, (*int32)(req.Msg.DefaultRole))
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&userv1.UpdateUserResponse{User: updated}), nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *connect.Request[userv1.DeleteUserRequest]) (*connect.Response[userv1.DeleteUserResponse], error) {
	if req.Msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("id is required"))
	}

	authCtx, err := s.requireAuth(ctx)
	if err != nil {
		return nil, err
	}

	if authCtx.SubjectID == req.Msg.GetId() {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("cannot delete your own user account"))
	}

	if err := s.repo.Delete(ctx, req.Msg.GetId()); err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&userv1.DeleteUserResponse{}), nil
}
