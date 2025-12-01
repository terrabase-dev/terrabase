package services

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/auth"
	authv1 "github.com/terrabase-dev/terrabase/specs/terrabase/auth/v1"
	authzv1 "github.com/terrabase-dev/terrabase/specs/terrabase/authz/v1"
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

func (a AuthAware) requireAdminOrSelf(authCtx *auth.Context, ownerType string, ownerID string) error {
	if authCtx == nil {
		return connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	if authCtx.HasScope(authzv1.Scope_SCOPE_ADMIN) {
		return nil
	}

	if ownerType == "user" && ownerID != "" && authCtx.SubjectID == ownerID {
		return nil
	}

	return connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
}

func scopesForRole(role userRolev1.UserRole) []authzv1.Scope {
	switch role {
	case userRolev1.UserRole_USER_ROLE_OWNER:
		return []authzv1.Scope{
			authzv1.Scope_SCOPE_ADMIN,
			authzv1.Scope_SCOPE_ORG_WRITE,
			authzv1.Scope_SCOPE_ORG_READ,
			authzv1.Scope_SCOPE_TEAM_WRITE,
			authzv1.Scope_SCOPE_TEAM_READ,
			authzv1.Scope_SCOPE_APPLICATION_WRITE,
			authzv1.Scope_SCOPE_APPLICATION_READ,
			authzv1.Scope_SCOPE_ENVIRONMENT_WRITE,
			authzv1.Scope_SCOPE_ENVIRONMENT_READ,
			authzv1.Scope_SCOPE_WORKSPACE_WRITE,
			authzv1.Scope_SCOPE_WORKSPACE_READ,
		}
	case userRolev1.UserRole_USER_ROLE_MAINTAINER:
		return []authzv1.Scope{
			authzv1.Scope_SCOPE_ORG_READ,
			authzv1.Scope_SCOPE_TEAM_WRITE,
			authzv1.Scope_SCOPE_TEAM_READ,
			authzv1.Scope_SCOPE_APPLICATION_WRITE,
			authzv1.Scope_SCOPE_APPLICATION_READ,
			authzv1.Scope_SCOPE_ENVIRONMENT_WRITE,
			authzv1.Scope_SCOPE_ENVIRONMENT_READ,
			authzv1.Scope_SCOPE_WORKSPACE_WRITE,
			authzv1.Scope_SCOPE_WORKSPACE_READ,
		}
	case userRolev1.UserRole_USER_ROLE_DEVELOPER:
		return []authzv1.Scope{
			authzv1.Scope_SCOPE_ORG_READ,
			authzv1.Scope_SCOPE_TEAM_READ,
			authzv1.Scope_SCOPE_APPLICATION_READ,
			authzv1.Scope_SCOPE_ENVIRONMENT_READ,
			authzv1.Scope_SCOPE_WORKSPACE_READ,
		}
	default:
		return []authzv1.Scope{
			authzv1.Scope_SCOPE_ORG_READ,
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
