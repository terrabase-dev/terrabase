package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/terrabase-dev/terrabase/internal/auth"
	"github.com/terrabase-dev/terrabase/internal/db"
	"github.com/terrabase-dev/terrabase/internal/repos"
	"github.com/terrabase-dev/terrabase/internal/rpcserver"
	"github.com/uptrace/bun"
)

func main() {
	logger := log.New(os.Stdout, "[bedrockd] ", log.LstdFlags|log.Lmicroseconds|log.LUTC)
	addr := resolveAddr()

	dbURL := os.Getenv("DATABASE_URL")
	bunDB, err := db.OpenPostgres(context.Background(), dbURL)
	if err != nil {
		logger.Fatalf("failed to connect to database: %v", err)
	}
	defer bunDB.Close()

	authenticator := buildAuthenticator(bunDB, logger)
	refreshPepper := os.Getenv("AUTH_REFRESH_TOKEN_PEPPER")
	if refreshPepper == "" {
		refreshPepper = os.Getenv("AUTH_API_KEY_PEPPER")
	}
	services := rpcserver.NewServicesWithAuth(bunDB, logger, authenticator.TokenVerifier(), refreshPepper)

	server := rpcserver.New(rpcserver.Config{
		Addr:          addr,
		Authenticator: authenticator,
	}, services, logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Printf("shutdown error: %v", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("server error: %v", err)
	}
}

func resolveAddr() string {
	if addr := os.Getenv("TERRABASE_RPC_ADDR"); addr != "" {
		return addr
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if strings.HasPrefix(port, ":") {
		return port
	}

	return ":" + port
}

func buildAuthenticator(db *bun.DB, logger *log.Logger) *auth.Authenticator {
	jwtSecret := os.Getenv("AUTH_JWT_SECRET")
	if jwtSecret == "" {
		logger.Fatalf("AUTH_JWT_SECRET is required for authentication")
	}

	jwtIssuer := os.Getenv("AUTH_JWT_ISSUER")
	jwtAudience := os.Getenv("AUTH_JWT_AUDIENCE")

	tokenVerifier, err := auth.NewTokenVerifier([]byte(jwtSecret), jwtIssuer, jwtAudience)
	if err != nil {
		logger.Fatalf("failed to configure token verifier: %v", err)
	}

	refreshPepper := os.Getenv("AUTH_REFRESH_TOKEN_PEPPER")
	if refreshPepper == "" {
		refreshPepper = os.Getenv("AUTH_API_KEY_PEPPER")
	}
	if refreshPepper == "" {
		logger.Fatalf("AUTH_REFRESH_TOKEN_PEPPER (or AUTH_API_KEY_PEPPER) is required for refresh token hashing")
	}

	apiKeyResolver := auth.NewAPIKeyResolver(
		repos.NewAPIKeyRepo(db),
		repos.NewUserRepo(db),
		os.Getenv("AUTH_API_KEY_PEPPER"),
	)

	return auth.NewAuthenticator(tokenVerifier, apiKeyResolver, repos.NewSessionRepo(db))
}
