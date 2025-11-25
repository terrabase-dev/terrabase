package rpcserver

import (
	"log"

	applicationv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/application/v1/applicationv1connect"
	authv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/auth/v1/authv1connect"
	driftReportv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/drift_report/v1/driftReportv1connect"
	environmentv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/environment/v1/environmentv1connect"
	lockv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/lock/v1/lockv1connect"
	organizationv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/organization/v1/organizationv1connect"
	s3BackendConfigv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/s3_backend_config/v1/s3BackendConfigv1connect"
	stateVersionv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/state_version/v1/stateVersionv1connect"
	teamv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/team/v1/teamv1connect"
	userv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/user/v1/userv1connect"
	userMembershipv1connect "github.com/terrabase-dev/terrabase/specs/terrabase/user_membership/v1/userMembershipv1connect"
	workspacev1connect "github.com/terrabase-dev/terrabase/specs/terrabase/workspace/v1/workspacev1connect"

	"github.com/terrabase-dev/terrabase/internal/auth"
	"github.com/terrabase-dev/terrabase/internal/repos"
	"github.com/terrabase-dev/terrabase/internal/services"
	"github.com/uptrace/bun"
)

type Services struct {
	Application     applicationv1connect.ApplicationServiceHandler
	Auth            authv1connect.AuthServiceHandler
	Environment     environmentv1connect.EnvironmentServiceHandler
	StateVersion    stateVersionv1connect.StateVersionServiceHandler
	User            userv1connect.UserServiceHandler
	Team            teamv1connect.TeamServiceHandler
	Workspace       workspacev1connect.WorkspaceServiceHandler
	Lock            lockv1connect.LockServiceHandler
	DriftReport     driftReportv1connect.DriftReportServiceHandler
	Organization    organizationv1connect.OrganizationServiceHandler
	UserMembership  userMembershipv1connect.UserMembershipServiceHandler
	S3BackendConfig s3BackendConfigv1connect.S3BackendConfigServiceHandler
}

func NewServices(db *bun.DB, logger *log.Logger) Services {
	return NewServicesWithAuth(db, logger, nil, "")
}

func NewServicesWithAuth(db *bun.DB, logger *log.Logger, tokenVerifier *auth.TokenVerifier, refreshPepper string) Services {
	return Services{
		Application:     applicationv1connect.UnimplementedApplicationServiceHandler{},
		Auth:            services.NewAuthService(repos.NewUserRepo(db), repos.NewCredentialRepo(db), repos.NewSessionRepo(db), repos.NewAPIKeyRepo(db), tokenVerifier, refreshPepper),
		Environment:     environmentv1connect.UnimplementedEnvironmentServiceHandler{},
		StateVersion:    stateVersionv1connect.UnimplementedStateVersionServiceHandler{},
		User:            services.NewUserService(repos.NewUserRepo(db), logger),
		Team:            services.NewTeamService(repos.NewTeamRepo(db), logger),
		Workspace:       workspacev1connect.UnimplementedWorkspaceServiceHandler{},
		Lock:            lockv1connect.UnimplementedLockServiceHandler{},
		DriftReport:     driftReportv1connect.UnimplementedDriftReportServiceHandler{},
		Organization:    services.NewOrganizationService(repos.NewOrganizationRepo(db), logger),
		UserMembership:  userMembershipv1connect.UnimplementedUserMembershipServiceHandler{},
		S3BackendConfig: s3BackendConfigv1connect.UnimplementedS3BackendConfigServiceHandler{},
	}
}
