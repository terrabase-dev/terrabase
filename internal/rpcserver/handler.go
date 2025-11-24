package rpcserver

import (
	"context"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/auth"
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
)

func buildHandler(services Services, authenticator *auth.Authenticator, logger *log.Logger) http.Handler {
	mux := http.NewServeMux()

	handlerOpts := []connect.HandlerOption{
		connect.WithInterceptors(auth.ContextInterceptor(authenticator, logger)),
		connect.WithInterceptors(loggingInterceptor(logger)),
	}

	path, handler := applicationv1connect.NewApplicationServiceHandler(services.Application, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = authv1connect.NewAuthServiceHandler(services.Auth, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = environmentv1connect.NewEnvironmentServiceHandler(services.Environment, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = stateVersionv1connect.NewStateVersionServiceHandler(services.StateVersion, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = userv1connect.NewUserServiceHandler(services.User, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = teamv1connect.NewTeamServiceHandler(services.Team, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = workspacev1connect.NewWorkspaceServiceHandler(services.Workspace, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = lockv1connect.NewLockServiceHandler(services.Lock, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = driftReportv1connect.NewDriftReportServiceHandler(services.DriftReport, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = organizationv1connect.NewOrganizationServiceHandler(services.Organization, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = userMembershipv1connect.NewUserMembershipServiceHandler(services.UserMembership, handlerOpts...)
	mux.Handle(path, handler)

	path, handler = s3BackendConfigv1connect.NewS3BackendConfigServiceHandler(services.S3BackendConfig, handlerOpts...)
	mux.Handle(path, handler)

	mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))

	return mux
}

func loggingInterceptor(logger *log.Logger) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()
			res, err := next(ctx, req)
			elapsed := time.Since(start).Round(time.Millisecond)
			if err != nil {
				logger.Printf("rpc %s error=%v duration=%s", req.Spec().Procedure, err, elapsed)
			} else {
				logger.Printf("rpc %s ok duration=%s", req.Spec().Procedure, elapsed)
			}
			return res, err
		}
	}
}
