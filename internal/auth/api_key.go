package auth

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/terrabase-dev/terrabase/internal/repos"
)

type APIKeyResolver struct {
	keys   *repos.APIKeyRepo
	users  *repos.UserRepo
	pepper []byte
	now    func() time.Time
}

func NewAPIKeyResolver(keys *repos.APIKeyRepo, users *repos.UserRepo, pepper string) *APIKeyResolver {
	return &APIKeyResolver{
		keys:   keys,
		users:  users,
		pepper: []byte(pepper),
		now:    time.Now,
	}
}

func (r *APIKeyResolver) Authenticate(ctx context.Context, rawKey string) (*Context, error) {
	if rawKey == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("api key is required"))
	}

	prefix, secret, ok := strings.Cut(rawKey, ".")
	if !ok || prefix == "" || secret == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid api key format"))
	}

	key, err := r.keys.GetActiveByPrefix(ctx, prefix)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	if !hmac.Equal([]byte(key.SecretHash), []byte(hashSecret(secret, r.pepper))) {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid api key"))
	}

	authCtx := &Context{
		SubjectID:     key.OwnerID,
		PrincipalType: mapOwnerType(key.OwnerType),
		Scopes:        key.Scopes,
		Metadata:      key.Metadata,
		TokenID:       key.ID,
		AuthScheme:    "api_key",
		Authenticated: true,
		RawCredential: rawKey,
	}

	if userID := key.OwnerID; authCtx.PrincipalType == PrincipalUser && r.users != nil {
		if user, getErr := r.users.Get(ctx, userID); getErr == nil {
			authCtx.Name = user.Name
			authCtx.Email = user.Email
			authCtx.DefaultRole = user.DefaultRole
		}
	}

	_ = r.keys.TouchLastUsed(ctx, key.ID, r.now().UTC())

	return authCtx, nil
}

func mapOwnerType(ownerType string) PrincipalType {
	switch strings.ToLower(ownerType) {
	case "user":
		return PrincipalUser
	case "bot":
		return PrincipalBot
	case "service":
		return PrincipalService
	default:
		return PrincipalUnknown
	}
}

func hashSecret(secret string, pepper []byte) string {
	h := hmac.New(sha256.New, pepper)
	_, _ = h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(nil))
}

func HashAPIKeySecret(secret string, pepper []byte) string {
	return hashSecret(secret, pepper)
}

func BuildAPIKey(prefix string, secret string) string {
	return fmt.Sprintf("%s.%s", prefix, secret)
}

// APIKeyMaterial holds generated parts of an API key. Call Token() to return
// the user-facing token and SecretHashWithPepper to store the hashed secret.
type APIKeyMaterial struct {
	Prefix string
	Secret string
}

func BuildAPIKeyMaterial() APIKeyMaterial {
	return APIKeyMaterial{
		Prefix: randomHex(6),
		Secret: randomSecret(32),
	}
}

func (m APIKeyMaterial) Token() string {
	return BuildAPIKey(m.Prefix, m.Secret)
}

func (m APIKeyMaterial) SecretHashWithPepper(pepper []byte) string {
	return hashSecret(m.Secret, pepper)
}

func randomHex(bytes int) string {
	b := make([]byte, bytes)
	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}

func randomSecret(bytes int) string {
	b := make([]byte, bytes)
	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}
