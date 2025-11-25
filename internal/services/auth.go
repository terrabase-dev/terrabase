package services

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/terrabase-dev/terrabase/internal/auth"
	"github.com/terrabase-dev/terrabase/internal/models"
	"github.com/terrabase-dev/terrabase/internal/repos"
	authv1 "github.com/terrabase-dev/terrabase/specs/terrabase/auth/v1"
	userv1 "github.com/terrabase-dev/terrabase/specs/terrabase/user/v1"
	userRolev1 "github.com/terrabase-dev/terrabase/specs/terrabase/user_role/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	defaultAccessTTL  = 15 * time.Minute
	defaultRefreshTTL = 24 * time.Hour
)

type AuthService struct {
	users       *repos.UserRepo
	creds       *repos.CredentialRepo
	sessions    *repos.SessionRepo
	apiKeys     *repos.APIKeyRepo
	tokenSigner *auth.TokenVerifier
	refreshPep  []byte
}

func NewAuthService(
	users *repos.UserRepo,
	creds *repos.CredentialRepo,
	sessions *repos.SessionRepo,
	apiKeys *repos.APIKeyRepo,
	tokenSigner *auth.TokenVerifier,
	refreshPepper string,
) *AuthService {
	return &AuthService{
		users:       users,
		creds:       creds,
		sessions:    sessions,
		apiKeys:     apiKeys,
		tokenSigner: tokenSigner,
		refreshPep:  []byte(refreshPepper),
	}
}

func (s *AuthService) Signup(ctx context.Context, req *connect.Request[authv1.SignupRequest]) (*connect.Response[authv1.SignupResponse], error) {
	if req.Msg.GetPassword() == "" || len(req.Msg.Password) < 8 {
		// TODO: add more password requirements
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("password must be at least 8 characters"))
	}

	user := &userv1.User{
		Name:        req.Msg.GetName(),
		Email:       req.Msg.GetEmail(),
		DefaultRole: req.Msg.GetDefaultRole(),
		UserType:    userv1.UserType_USER_TYPE_USER,
	}

	created, err := s.users.Create(ctx, user)
	if err != nil {
		return nil, mapError(err)
	}

	hash, err := auth.HashPassword(req.Msg.Password)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("hash password: %w", err))
	}

	if err := s.creds.UpsertPassword(ctx, created.GetId(), hash, "argon2id"); err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("store credential: %w", err))
	}

	accessToken, refreshToken, err := s.issueTokens(ctx, created, req)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&authv1.SignupResponse{
		User:         created,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}), nil
}

func (s *AuthService) Login(ctx context.Context, req *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	user, err := s.users.GetByEmail(ctx, req.Msg.GetEmail())
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid credentials"))
	}

	cred, err := s.creds.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid credentials"))
	}

	ok, err := auth.VerifyPassword(req.Msg.GetPassword(), cred.PasswordHash)
	if err != nil || !ok {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid credentials"))
	}

	accessToken, refreshToken, err := s.issueTokens(ctx, user.ToProto(), req)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&authv1.LoginResponse{
		User:         user.ToProto(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}), nil
}

func (s *AuthService) Refresh(ctx context.Context, req *connect.Request[authv1.RefreshRequest]) (*connect.Response[authv1.RefreshResponse], error) {
	refresh := req.Msg.GetRefreshToken()
	if refresh == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("refresh token required"))
	}

	hash := s.hashRefreshToken(refresh)

	session, err := s.sessions.GetByRefreshHash(ctx, hash)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid refresh token"))
	}

	user, err := s.users.Get(ctx, session.UserID)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid refresh token"))
	}

	accessToken, newRefreshToken, err := s.issueTokens(ctx, user.ToProto(), req)
	if err != nil {
		return nil, err
	}

	// Rotate: delete old session
	_ = s.sessions.DeleteByID(ctx, session.ID)

	return connect.NewResponse(&authv1.RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}), nil
}

func (s *AuthService) WhoAmI(ctx context.Context, _ *connect.Request[authv1.WhoAmIRequest]) (*connect.Response[authv1.WhoAmIResponse], error) {
	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	user, err := s.users.Get(ctx, authCtx.SubjectID)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	return connect.NewResponse(&authv1.WhoAmIResponse{
		User:   user.ToProto(),
		Scopes: authCtx.Scopes,
	}), nil
}

func (s *AuthService) Logout(ctx context.Context, req *connect.Request[authv1.LogoutRequest]) (*connect.Response[authv1.LogoutResponse], error) {
	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	targetSessionID := req.Msg.GetSessionId()
	if targetSessionID == "" {
		targetSessionID = authCtx.TokenID
	}

	if targetSessionID == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("session_id is required"))
	}

	if err := s.sessions.DeleteByID(ctx, targetSessionID); err != nil && !errors.Is(err, repos.ErrNotFound) {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&authv1.LogoutResponse{}), nil
}

func (s *AuthService) ListSessions(ctx context.Context, _ *connect.Request[authv1.ListSessionsRequest]) (*connect.Response[authv1.ListSessionsResponse], error) {
	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	sessions, err := s.sessions.ListByUser(ctx, authCtx.SubjectID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp := &authv1.ListSessionsResponse{Sessions: make([]*authv1.Session, 0, len(sessions))}
	for i := range sessions {
		s := sessions[i]
		resp.Sessions = append(resp.Sessions, &authv1.Session{
			Id:         s.ID,
			UserAgent:  s.UserAgent,
			Ip:         s.IP,
			ExpiresAt:  timestamppb.New(s.ExpiresAt.UTC()),
			CreatedAt:  timestamppb.New(s.CreatedAt.UTC()),
			LastUsedAt: timestamppb.New(s.LastUsedAt.UTC()),
		})
	}

	return connect.NewResponse(resp), nil
}

func (s *AuthService) CreateMachineUser(ctx context.Context, req *connect.Request[authv1.CreateMachineUserRequest]) (*connect.Response[authv1.CreateMachineUserResponse], error) {
	if req.Msg.GetUserType() == userv1.UserType_USER_TYPE_USER {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("user_type must be one of \"bot\" or \"service\""))
	}

	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	if authCtx == nil || !authCtx.HasScope(authv1.Scope_SCOPE_ADMIN) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	ownerID := req.Msg.GetOwnerUserId()
	if ownerID == "" {
		ownerID = authCtx.SubjectID
	}

	user := &userv1.User{
		Name:        req.Msg.GetName(),
		DefaultRole: req.Msg.GetDefaultRole(),
		UserType:    req.Msg.GetUserType(),
		OwnerUserId: &ownerID,
	}

	created, err := s.users.Create(ctx, user)
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&authv1.CreateMachineUserResponse{
		MachineUser: created,
	}), nil
}

func (s *AuthService) issueTokens(ctx context.Context, user *userv1.User, req connect.AnyRequest) (accessToken string, refreshToken string, err error) {
	if s.tokenSigner == nil {
		return "", "", connect.NewError(connect.CodeInternal, errors.New("token signer not configured"))
	}

	scopes := scopesForRole(user.GetDefaultRole())

	sessionID := uuid.NewString()
	refreshToken = uuid.NewString() + "." + uuid.NewString()
	refreshHash := s.hashRefreshToken(refreshToken)

	ua, ip := extractClientInfo(req)

	if _, err := s.sessions.Create(ctx, sessionID, user.GetId(), refreshHash, time.Now().Add(defaultRefreshTTL), ua, ip, map[string]any{}); err != nil {
		return "", "", connect.NewError(connect.CodeInternal, fmt.Errorf("create session: %w", err))
	}

	claims := &auth.Claims{
		SubjectType:  string(auth.PrincipalUser),
		Name:         user.GetName(),
		Email:        user.GetEmail(),
		Metadata:     map[string]any{},
		Entitlements: map[string][]string{},
		Scopes:       scopes,
		RegisteredClaims: auth.RegisteredClaimsForTTLWithID(
			user.GetId(),
			sessionID,
			defaultAccessTTL,
		),
	}

	accessToken, err = s.tokenSigner.Issue(claims)
	if err != nil {
		return "", "", connect.NewError(connect.CodeInternal, fmt.Errorf("issue access token: %w", err))
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) hashRefreshToken(token string) string {
	mac := hmac.New(sha256.New, s.refreshPep)
	_, _ = mac.Write([]byte(token))

	return hex.EncodeToString(mac.Sum(nil))
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

func (s *AuthService) CreateApiKey(ctx context.Context, req *connect.Request[authv1.CreateApiKeyRequest]) (*connect.Response[authv1.CreateApiKeyResponse], error) {
	if s.apiKeys == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("api key store not configured"))
	}

	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	ownerType := apiKeyOwnerTypeToString(req.Msg.GetOwnerType())
	if ownerType == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("owner_type is required"))
	}

	ownerID := req.Msg.GetOwnerId()
	if ownerType == "user" && ownerID == "" {
		ownerID = authCtx.SubjectID
	}

	if !isAdminOrSelfUser(authCtx, ownerType, ownerID) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	if len(req.Msg.Scopes) == 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("scopes are required"))
	}

	exp := time.Time{}

	if req.Msg.TtlHours != nil && *req.Msg.TtlHours > 0 {
		exp = time.Now().Add(time.Duration(*req.Msg.TtlHours) * time.Hour)
	}

	mat := auth.BuildAPIKeyMaterial()

	key := &models.APIKey{
		ID:         uuid.NewString(),
		OwnerType:  ownerType,
		OwnerID:    ownerID,
		Name:       req.Msg.GetName(),
		Prefix:     mat.Prefix,
		SecretHash: mat.SecretHashWithPepper(s.refreshPep),
		Scopes:     req.Msg.Scopes,
		ExpiresAt:  exp,
	}

	created, err := s.apiKeys.Create(ctx, key)
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(&authv1.CreateApiKeyResponse{
		ApiKeyToken: mat.Token(),
		ApiKey:      toProtoAPIKey(created),
	}), nil
}

func (s *AuthService) ListApiKeys(ctx context.Context, req *connect.Request[authv1.ListApiKeysRequest]) (*connect.Response[authv1.ListApiKeysResponse], error) {
	if s.apiKeys == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("api key store not configured"))
	}

	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	ownerType := apiKeyOwnerTypeToString(req.Msg.GetOwnerType())
	ownerID := req.Msg.GetOwnerId()

	if ownerType == "user" && ownerID == "" {
		ownerID = authCtx.SubjectID
	}

	if !isAdminOrSelfUser(authCtx, ownerType, ownerID) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	keys, err := s.apiKeys.ListByOwner(ctx, ownerType, ownerID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp := &authv1.ListApiKeysResponse{ApiKeys: make([]*authv1.ApiKey, 0, len(keys))}

	for _, k := range keys {
		resp.ApiKeys = append(resp.ApiKeys, toProtoAPIKey(k))
	}

	return connect.NewResponse(resp), nil
}

func (s *AuthService) RevokeApiKey(ctx context.Context, req *connect.Request[authv1.RevokeApiKeyRequest]) (*connect.Response[authv1.RevokeApiKeyResponse], error) {
	if s.apiKeys == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("api key store not configured"))
	}

	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	existing, err := s.apiKeys.GetByID(ctx, req.Msg.GetId())
	if err != nil {
		return nil, mapError(err)
	}

	if !isAdminOrSelfUser(authCtx, existing.OwnerType, existing.OwnerID) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	if err := s.apiKeys.Revoke(ctx, req.Msg.GetId(), req.Msg.GetReason()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&authv1.RevokeApiKeyResponse{}), nil
}

func (s *AuthService) RotateApiKey(ctx context.Context, req *connect.Request[authv1.RotateApiKeyRequest]) (*connect.Response[authv1.RotateApiKeyResponse], error) {
	if s.apiKeys == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("api key store not configured"))
	}

	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	existing, err := s.apiKeys.GetByID(ctx, req.Msg.GetId())
	if err != nil {
		return nil, mapError(err)
	}

	if !isAdminOrSelfUser(authCtx, existing.OwnerType, existing.OwnerID) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	scopes := req.Msg.GetScopes()

	if len(scopes) == 0 {
		scopes = existing.Scopes
	}

	exp := existing.ExpiresAt

	if req.Msg.TtlHours != nil && *req.Msg.TtlHours > 0 {
		exp = time.Now().Add(time.Duration(*req.Msg.TtlHours) * time.Hour)
	}

	mat := auth.BuildAPIKeyMaterial()

	newKey := &models.APIKey{
		ID:          uuid.NewString(),
		OwnerType:   existing.OwnerType,
		OwnerID:     existing.OwnerID,
		Name:        existing.Name,
		Prefix:      mat.Prefix,
		SecretHash:  mat.SecretHashWithPepper(s.refreshPep),
		Scopes:      scopes,
		ExpiresAt:   exp,
		RotatedFrom: existing.ID,
		Metadata:    existing.Metadata,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if _, err := s.apiKeys.Create(ctx, newKey); err != nil {
		return nil, mapError(err)
	}

	_ = s.apiKeys.Revoke(ctx, existing.ID, "rotated")

	return connect.NewResponse(&authv1.RotateApiKeyResponse{
		ApiKeyToken: mat.Token(),
		ApiKey:      toProtoAPIKey(newKey),
	}), nil
}

func (s *AuthService) requireScope(ctx context.Context, scope authv1.Scope) error {
	authCtx, ok := auth.FromContext(ctx)
	if !ok || !authCtx.Authenticated {
		return connect.NewError(connect.CodeUnauthenticated, errors.New("unauthenticated"))
	}

	if !authCtx.HasScope(scope) {
		return connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	return nil
}

func toProtoAPIKey(key *models.APIKey) *authv1.ApiKey {
	if key == nil {
		return nil
	}

	resp := &authv1.ApiKey{
		Id:        key.ID,
		Name:      key.Name,
		Scopes:    key.Scopes,
		OwnerId:   key.OwnerID,
		OwnerType: toProtoOwnerType(key.OwnerType),
		CreatedAt: timestamppb.New(key.CreatedAt.UTC()),
	}

	if !key.ExpiresAt.IsZero() {
		resp.ExpiresAt = timestamppb.New(key.ExpiresAt.UTC())
	}

	if !key.LastUsedAt.IsZero() {
		resp.LastUsedAt = timestamppb.New(key.LastUsedAt.UTC())
	}

	if !key.RevokedAt.IsZero() {
		resp.RevokedAt = timestamppb.New(key.RevokedAt.UTC())
	}

	return resp
}

func toProtoOwnerType(ownerType string) authv1.ApiKeyOwnerType {
	switch ownerType {
	case "user":
		return authv1.ApiKeyOwnerType_API_KEY_OWNER_TYPE_USER
	case "bot":
		return authv1.ApiKeyOwnerType_API_KEY_OWNER_TYPE_BOT
	case "service":
		return authv1.ApiKeyOwnerType_API_KEY_OWNER_TYPE_SERVICE
	default:
		return authv1.ApiKeyOwnerType_API_KEY_OWNER_TYPE_UNSPECIFIED
	}
}

func extractClientInfo(req connect.AnyRequest) (userAgent string, ip string) {
	if req == nil {
		return "", ""
	}

	userAgent = req.Header().Get("user-agent")

	xff := req.Header().Get("x-forwarded-for")
	if xff != "" {
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			ip = strings.TrimSpace(parts[0])
		}
	}

	if ip == "" {
		ip = req.Header().Get("x-real-ip")
	}

	if ip == "" {
		p := req.Peer()
		ip = p.Addr
	}

	return userAgent, ip
}

func isAdminOrSelfUser(authCtx *auth.Context, ownerType string, ownerID string) bool {
	if authCtx == nil {
		return false
	}

	if authCtx.HasScope(authv1.Scope_SCOPE_ADMIN) {
		return true
	}

	if ownerType == "user" && ownerID != "" && authCtx.SubjectID == ownerID {
		return true
	}

	return false
}
