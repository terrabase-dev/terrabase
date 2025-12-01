package rpcserver

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/auth"
	applicationv1 "github.com/terrabase-dev/terrabase/specs/terrabase/application/v1"
	authv1 "github.com/terrabase-dev/terrabase/specs/terrabase/auth/v1"
	authzv1 "github.com/terrabase-dev/terrabase/specs/terrabase/authz/v1"
	driftReportv1 "github.com/terrabase-dev/terrabase/specs/terrabase/drift_report/v1"
	environmentv1 "github.com/terrabase-dev/terrabase/specs/terrabase/environment/v1"
	lockv1 "github.com/terrabase-dev/terrabase/specs/terrabase/lock/v1"
	organizationv1 "github.com/terrabase-dev/terrabase/specs/terrabase/organization/v1"
	s3BackendConfigv1 "github.com/terrabase-dev/terrabase/specs/terrabase/s3_backend_config/v1"
	stateVersionv1 "github.com/terrabase-dev/terrabase/specs/terrabase/state_version/v1"
	teamv1 "github.com/terrabase-dev/terrabase/specs/terrabase/team/v1"
	userv1 "github.com/terrabase-dev/terrabase/specs/terrabase/user/v1"
	userMembershipv1 "github.com/terrabase-dev/terrabase/specs/terrabase/user_membership/v1"
	workspacev1 "github.com/terrabase-dev/terrabase/specs/terrabase/workspace/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type methodPolicy struct {
	authRequired bool
	scopes       []authzv1.Scope
	adminOrSelf  bool
	ownerIdField string
}

// buildMethodPolicies reads custom auth options from proto descriptors and builds a lookup by procedure.
func buildMethodPolicies() map[string]methodPolicy {
	files := []protoreflect.FileDescriptor{
		authv1.File_terrabase_auth_v1_auth_proto,
		applicationv1.File_terrabase_application_v1_application_proto,
		driftReportv1.File_terrabase_drift_report_v1_drift_report_proto,
		environmentv1.File_terrabase_environment_v1_environment_proto,
		lockv1.File_terrabase_lock_v1_lock_proto,
		organizationv1.File_terrabase_organization_v1_organization_proto,
		s3BackendConfigv1.File_terrabase_s3_backend_config_v1_s3_backend_config_proto,
		stateVersionv1.File_terrabase_state_version_v1_state_version_proto,
		teamv1.File_terrabase_team_v1_team_proto,
		userMembershipv1.File_terrabase_user_membership_v1_user_membership_proto,
		userv1.File_terrabase_user_v1_user_proto,
		workspacev1.File_terrabase_workspace_v1_workspace_proto,
	}

	policies := make(map[string]methodPolicy)

	for _, fd := range files {
		for i := 0; i < fd.Services().Len(); i++ {
			svc := fd.Services().Get(i)

			for j := 0; j < svc.Methods().Len(); j++ {
				m := svc.Methods().Get(j)
				opts, _ := m.Options().(*descriptorpb.MethodOptions)
				policies[procedureName(fd.Package(), svc.Name(), m.Name())] = extractPolicy(opts)
			}
		}
	}

	return policies
}

func extractPolicy(opts *descriptorpb.MethodOptions) methodPolicy {
	var pol methodPolicy

	if opts == nil {
		return pol
	}

	if proto.HasExtension(opts, authzv1.E_AuthRequired) {
		if v, ok := proto.GetExtension(opts, authzv1.E_AuthRequired).(bool); ok {
			pol.authRequired = v
		}
	}

	if proto.HasExtension(opts, authzv1.E_RequiredScopes) {
		if vals, ok := proto.GetExtension(opts, authzv1.E_RequiredScopes).([]authzv1.Scope); ok {
			pol.scopes = append(pol.scopes, vals...)
		}
	}

	if proto.HasExtension(opts, authzv1.E_AdminOrSelf) {
		if v, ok := proto.GetExtension(opts, authzv1.E_AdminOrSelf).(bool); ok {
			pol.adminOrSelf = v
		}
	}

	if proto.HasExtension(opts, authzv1.E_OwnerIdField) {
		if v, ok := proto.GetExtension(opts, authzv1.E_OwnerIdField).(string); ok {
			pol.ownerIdField = v
		}
	}

	return pol
}

func procedureName(pkg protoreflect.FullName, svc protoreflect.Name, method protoreflect.Name) string {
	return fmt.Sprintf("/%s.%s/%s", pkg, svc, method)
}

func extractStringField(msg proto.Message, ownerIdField string) (string, error) {
	if msg == nil {
		return "", errors.New("missing request message")
	}

	v := msg.ProtoReflect()
	fd := v.Descriptor().Fields().ByName(protoreflect.Name(ownerIdField))
	if fd == nil {
		return "", fmt.Errorf("owner id field %q not found", ownerIdField)
	}

	if fd.Kind() != protoreflect.StringKind {
		return "", fmt.Errorf("owner id field %q must be string", ownerIdField)
	}

	return v.Get(fd).String(), nil
}

func authzInterceptor(policies map[string]methodPolicy) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			pol, ok := policies[req.Spec().Procedure]
			if !ok || !pol.authRequired {
				return next(ctx, req)
			}

			authCtx, ok := auth.FromContext(ctx)
			if !ok || !authCtx.Authenticated {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
			}

			if pol.adminOrSelf {
				if authCtx.HasScope(authzv1.Scope_SCOPE_ADMIN) {
					return next(ctx, req)
				}

				if pol.ownerIdField == "" {
					// WARNING: If we don't have the owner ID field, the handler will be required to enforce the admin or self requirement
					// TODO: we could probably wire a custom resolver interface for these
					return next(ctx, req)
				}

				msg, ok := req.Any().(proto.Message)
				if !ok {
					return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid request type"))
				}

				ownerId, err := extractStringField(msg, pol.ownerIdField)
				if err != nil {
					return nil, connect.NewError(connect.CodeInvalidArgument, err)
				}

				if authCtx.SubjectID != ownerId {
					return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
				}

				return next(ctx, req)
			}

			if len(pol.scopes) > 0 {
				if slices.ContainsFunc(pol.scopes, authCtx.HasScope) {
					return next(ctx, req)
				}

				return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
			}

			return next(ctx, req)
		}
	}
}
