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

	"github.com/terrabase-dev/terrabase/internal/db"
	"github.com/terrabase-dev/terrabase/internal/rpcserver"
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

	services := rpcserver.NewServices(bunDB, logger)

	server := rpcserver.New(rpcserver.Config{Addr: addr}, services, logger)

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
