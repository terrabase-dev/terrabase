package services

import (
	"context"

	"github.com/terrabase-dev/terrabase/internal/auth"
)

// AuthContext returns the authentication context attached by the Connect interceptor.
func AuthContext(ctx context.Context) (*auth.Context, bool) {
	return auth.FromContext(ctx)
}
