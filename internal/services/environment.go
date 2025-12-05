package services

import (
	"context"
	"log"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
	environmentv1 "github.com/terrabase-dev/terrabase/specs/terrabase/environment/v1"
)

type EnvironmentService struct {
	repo   *repos.EnvironmentRepo
	logger *log.Logger
}

func NewEnvironmentService(repo *repos.EnvironmentRepo, logger *log.Logger) *EnvironmentService {
	return &EnvironmentService{
		repo:   repo,
		logger: logger,
	}
}

func (s *EnvironmentService) CreateEnvironment(ctx context.Context, req *connect.Request[environmentv1.CreateEnvironmentRequest]) (*connect.Response[environmentv1.CreateEnvironmentResponse], error) {

}
