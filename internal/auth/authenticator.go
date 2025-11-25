package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
)

type Authenticator struct {
	tokens   *TokenVerifier
	apiKeys  *APIKeyResolver
	sessions *repos.SessionRepo
}

func NewAuthenticator(tokens *TokenVerifier, apiKeys *APIKeyResolver, sessions *repos.SessionRepo) *Authenticator {
	return &Authenticator{
		tokens:   tokens,
		apiKeys:  apiKeys,
		sessions: sessions,
	}
}

func (a *Authenticator) TokenVerifier() *TokenVerifier {
	if a == nil {
		return nil
	}

	return a.tokens
}

// Authenticate inspects the inbound headers and returns an auth context if credentials are present.
func (a *Authenticator) Authenticate(ctx context.Context, authHeader string, apiKeyHeader string) (*Context, error) {
	authorization := strings.TrimSpace(authHeader)

	if authorization != "" {
		switch {
		case strings.HasPrefix(strings.ToLower(authorization), "bearer "):
			token := strings.TrimSpace(authorization[len("bearer "):])

			if a.tokens == nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("token verifier not configured"))
			}

			authCtx, err := a.tokens.Verify(ctx, token)
			if err != nil {
				return nil, err
			}

			authCtx.RawCredential = authorization
			a.touchSession(ctx, authCtx)

			return authCtx, nil
		case strings.HasPrefix(strings.ToLower(authorization), "apikey "):
			apiKey := strings.TrimSpace(authorization[len("apikey "):])

			if a.apiKeys == nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("api key verifier not configured"))
			}

			authCtx, err := a.apiKeys.Authenticate(ctx, apiKey)
			if err != nil {
				return nil, err
			}

			authCtx.RawCredential = authorization

			return authCtx, nil
		default:
			return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unsupported authorization scheme"))
		}
	}

	apiKeyHeader = strings.TrimSpace(apiKeyHeader)
	if apiKeyHeader != "" {
		if a.apiKeys == nil {
			return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("api key verifier not configured"))
		}

		authCtx, err := a.apiKeys.Authenticate(ctx, apiKeyHeader)
		if err != nil {
			return nil, err
		}

		authCtx.RawCredential = apiKeyHeader

		return authCtx, nil
	}

	return nil, nil
}

// ContextInterceptor injects the authentication context into request context when credentials are present.
func ContextInterceptor(a *Authenticator, logger *log.Logger) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if a == nil {
				return next(ctx, req)
			}

			authCtx, err := a.Authenticate(ctx, req.Header().Get("Authorization"), req.Header().Get("X-API-Key"))
			if err != nil {
				return nil, err
			}

			if authCtx != nil {
				ctx = WithContext(ctx, authCtx)
			}

			res, callErr := next(ctx, req)

			return res, callErr
		}
	}
}

// AttachAuthHeaders returns a shallow copy of headers with an Authorization or X-API-Key
// suitable for outbound Connect clients.
func AttachAuthHeaders(headers http.Header, authCtx *Context) http.Header {
	if authCtx == nil || authCtx.RawCredential == "" {
		return headers
	}

	cloned := headers.Clone()
	if strings.HasPrefix(strings.ToLower(authCtx.RawCredential), "bearer ") || strings.HasPrefix(strings.ToLower(authCtx.RawCredential), "apikey ") {
		cloned.Set("Authorization", authCtx.RawCredential)
	} else {
		cloned.Set("X-API-Key", authCtx.RawCredential)
	}

	return cloned
}

func (a *Authenticator) touchSession(ctx context.Context, authCtx *Context) {
	if a.sessions == nil || authCtx == nil {
		return
	}

	if authCtx.AuthScheme != "access_token" {
		return
	}

	if authCtx.TokenID == "" {
		return
	}

	_ = a.sessions.TouchLastUsed(ctx, authCtx.TokenID, time.Now().UTC())
}
