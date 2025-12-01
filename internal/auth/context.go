package auth

import (
	"context"
	"slices"

	authzv1 "github.com/terrabase-dev/terrabase/specs/terrabase/authz/v1"
)

type PrincipalType string

const (
	PrincipalUser    PrincipalType = "user"
	PrincipalService PrincipalType = "service"
	PrincipalBot     PrincipalType = "bot"
	PrincipalUnknown PrincipalType = "unknown"
)

type Context struct {
	SubjectID     string
	PrincipalType PrincipalType
	Name          string
	Email         string
	DefaultRole   int32
	Scopes        []authzv1.Scope
	Entitlements  map[string][]string
	Metadata      map[string]any
	TokenID       string
	AuthScheme    string // access_token | api_key
	RawCredential string // the original Authorization/API key header value
	Authenticated bool
}

func (c *Context) HasScope(scope authzv1.Scope) bool {
	return slices.Contains(c.Scopes, scope)
}

type ctxKey struct{}

func WithContext(ctx context.Context, authCtx *Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, authCtx)
}

func FromContext(ctx context.Context) (*Context, bool) {
	val := ctx.Value(ctxKey{})
	if val == nil {
		return nil, false
	}

	if authCtx, ok := val.(*Context); ok {
		return authCtx, true
	}

	return nil, false
}
