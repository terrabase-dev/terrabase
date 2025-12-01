package auth

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	authzv1 "github.com/terrabase-dev/terrabase/specs/terrabase/authz/v1"
)

type TokenVerifier struct {
	secret   []byte
	issuer   string
	audience string
	now      func() time.Time
}

type Claims struct {
	SubjectType  string              `json:"sub_type,omitempty"`
	Name         string              `json:"name,omitempty"`
	Email        string              `json:"email,omitempty"`
	Scopes       []authzv1.Scope     `json:"scopes,omitempty"`
	Metadata     map[string]any      `json:"metadata,omitempty"`
	Entitlements map[string][]string `json:"entitlements,omitempty"`
	jwt.RegisteredClaims
}

func NewTokenVerifier(secret []byte, issuer string, audience string) (*TokenVerifier, error) {
	if len(secret) == 0 {
		return nil, errors.New("token verifier secret is required")
	}

	return &TokenVerifier{
		secret:   secret,
		issuer:   issuer,
		audience: audience,
		now:      time.Now,
	}, nil
}

func (v *TokenVerifier) Verify(_ context.Context, token string) (*Context, error) {
	claims := &Claims{}

	parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Method.Alg())
		}

		return v.secret, nil
	})

	if err != nil || !parsed.Valid {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid token: %w", err))
	}

	if v.issuer != "" && claims.Issuer != v.issuer {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("issuer mismatch"))
	}

	if v.audience != "" {
		matched := slices.Contains(claims.Audience, v.audience)
		if !matched {
			return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("audience mismatch"))
		}
	}

	now := v.now()
	if claims.ExpiresAt != nil && now.After(claims.ExpiresAt.Time.Add(30*time.Second)) {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("token expired"))
	}

	if claims.NotBefore != nil && now.Before(claims.NotBefore.Time.Add(-30*time.Second)) {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("token not yet valid"))
	}

	if claims.IssuedAt != nil && now.Before(claims.IssuedAt.Time.Add(-30*time.Second)) {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("token used before issued"))
	}

	principalType := PrincipalType(claims.SubjectType)
	if principalType == "" {
		principalType = PrincipalUser
	}

	authCtx := &Context{
		SubjectID:     claims.Subject,
		PrincipalType: principalType,
		Name:          claims.Name,
		Email:         claims.Email,
		Scopes:        claims.Scopes,
		Entitlements:  claims.Entitlements,
		Metadata:      claims.Metadata,
		TokenID:       claims.ID,
		AuthScheme:    "access_token",
		Authenticated: true,
	}

	return authCtx, nil
}

// Issue signs new JWT access tokens using the configured secret.
func (v *TokenVerifier) Issue(claims *Claims) (string, error) {
	if claims == nil {
		return "", errors.New("claims required")
	}

	if claims.Issuer == "" && v.issuer != "" {
		claims.Issuer = v.issuer
	}

	if len(claims.Audience) == 0 && v.audience != "" {
		claims.Audience = []string{v.audience}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(v.secret)
}

// RegisteredClaimsForTTL returns RegisteredClaims populated with iss/aud/iat/nbf/exp/jti.
func RegisteredClaimsForTTL(subject string, ttl time.Duration) jwt.RegisteredClaims {
	now := time.Now()

	return jwt.RegisteredClaims{
		Subject:   subject,
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		ID:        uuid.NewString(),
	}
}

// RegisteredClaimsForTTLWithID allows specifying a particular token ID (e.g., session ID).
func RegisteredClaimsForTTLWithID(subject string, id string, ttl time.Duration) jwt.RegisteredClaims {
	now := time.Now()

	if id == "" {
		id = uuid.NewString()
	}

	return jwt.RegisteredClaims{
		Subject:   subject,
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		ID:        id,
	}
}
