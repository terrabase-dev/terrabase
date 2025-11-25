package services

import (
	"context"
	"errors"
	"slices"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/auth"
	authv1 "github.com/terrabase-dev/terrabase/specs/terrabase/auth/v1"
	userRolev1 "github.com/terrabase-dev/terrabase/specs/terrabase/user_role/v1"
)

type AuthAware struct{}

func (AuthAware) requireAuth(ctx context.Context) (*auth.Context, error) {
	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	return authCtx, nil
}

func (a AuthAware) requireScope(ctx context.Context, scope authv1.Scope) (*auth.Context, error) {
	authCtx, err := a.requireAuth(ctx)

	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	if !authCtx.HasScope(scope) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	return authCtx, nil
}

func (a AuthAware) requireAnyScope(ctx context.Context, scopes ...authv1.Scope) (*auth.Context, error) {
	authCtx, err := a.requireAuth(ctx)

	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	if len(scopes) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("scopes are required"))
	}

	if slices.ContainsFunc(scopes, authCtx.HasScope) {
		return authCtx, nil
	}

	return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
}

func (a AuthAware) requireAdminOrSelf(authCtx *auth.Context, ownerType string, ownerID string) error {
	if authCtx == nil {
		return connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	if authCtx.HasScope(authv1.Scope_SCOPE_ADMIN) {
		return nil
	}

	if ownerType == "user" && ownerID != "" && authCtx.SubjectID == ownerID {
		return nil
	}

	return connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
}

func scopesForRole(role userRolev1.UserRole) []authv1.Scope {
	switch role {
	case userRolev1.UserRole_USER_ROLE_OWNER:
		return []authv1.Scope{
			authv1.Scope_SCOPE_ADMIN,
			authv1.Scope_SCOPE_ORG_WRITE,
			authv1.Scope_SCOPE_ORG_READ,
			authv1.Scope_SCOPE_TEAM_WRITE,
			authv1.Scope_SCOPE_TEAM_READ,
			authv1.Scope_SCOPE_APPLICATION_WRITE,
			authv1.Scope_SCOPE_APPLICATION_READ,
			authv1.Scope_SCOPE_ENVIRONMENT_WRITE,
			authv1.Scope_SCOPE_ENVIRONMENT_READ,
			authv1.Scope_SCOPE_WORKSPACE_WRITE,
			authv1.Scope_SCOPE_WORKSPACE_READ,
		}
	case userRolev1.UserRole_USER_ROLE_MAINTAINER:
		return []authv1.Scope{
			authv1.Scope_SCOPE_ORG_READ,
			authv1.Scope_SCOPE_TEAM_WRITE,
			authv1.Scope_SCOPE_TEAM_READ,
			authv1.Scope_SCOPE_APPLICATION_WRITE,
			authv1.Scope_SCOPE_APPLICATION_READ,
			authv1.Scope_SCOPE_ENVIRONMENT_WRITE,
			authv1.Scope_SCOPE_ENVIRONMENT_READ,
			authv1.Scope_SCOPE_WORKSPACE_WRITE,
			authv1.Scope_SCOPE_WORKSPACE_READ,
		}
	case userRolev1.UserRole_USER_ROLE_DEVELOPER:
		return []authv1.Scope{
			authv1.Scope_SCOPE_ORG_READ,
			authv1.Scope_SCOPE_TEAM_READ,
			authv1.Scope_SCOPE_APPLICATION_READ,
			authv1.Scope_SCOPE_ENVIRONMENT_READ,
			authv1.Scope_SCOPE_WORKSPACE_READ,
		}
	default:
		return []authv1.Scope{
			authv1.Scope_SCOPE_ORG_READ,
		}
	}
}

func apiKeyOwnerTypeToString(ownerType authv1.ApiKeyOwnerType) string {
	switch ownerType {
	case authv1.ApiKeyOwnerType_API_KEY_OWNER_TYPE_USER:
		return "user"
	case authv1.ApiKeyOwnerType_API_KEY_OWNER_TYPE_BOT:
		return "bot"
	case authv1.ApiKeyOwnerType_API_KEY_OWNER_TYPE_SERVICE:
		return "service"
	default:
		return ""
	}
}
