package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/terrabase-dev/terrabase/internal/db"
)

func main() {
	logger := log.New(os.Stdout, "[bedrockmigrate] ", log.LstdFlags|log.Lmicroseconds|log.LUTC)

	dbURL := os.Getenv("DATABASE_URL")
	bunDB, err := db.OpenPostgres(context.Background(), dbURL)
	if err != nil {
		logger.Fatalf("failed to connect to database: %v", err)
	}
	defer bunDB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := db.RunMigrations(ctx, bunDB, logger); err != nil {
		logger.Fatalf("failed to run migrations: %v", err)
	}

	logger.Printf("migrations complete")
}
