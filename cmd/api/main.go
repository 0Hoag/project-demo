package main

import (
	"context"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zeross/project-demo/internal/httpserver"
	"github.com/zeross/project-demo/pkg/encrypter"
	"github.com/zeross/project-demo/pkg/log"
)

func main() {
	ctx := context.Background()

	port := 8080
	if v := os.Getenv("PORT"); v != "" {
		if p, err := strconv.Atoi(v); err == nil {
			port = p
		}
	}

	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "development"
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		jwtSecret = "dev-secret"
	}

	internalKey := os.Getenv("INTERNAL_KEY")
	if internalKey == "" {
		internalKey = "dev-internal-key"
	}

	l := log.InitializeZapLogger(log.ZapConfig{
		Level:    "debug",
		Mode:     mode,
		Encoding: "console",
	})

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		l.Fatalf(ctx, "DATABASE_URL is required")
	}

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		l.Fatalf(ctx, "pgxpool.New: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		l.Fatalf(ctx, "db.Ping: %v", err)
	}
	l.Infof(ctx, "Connected to Postgres successfully")

	srv := httpserver.New(l, httpserver.Config{
		Port:         port,
		DB:           pool,
		Mode:         mode,
		JwtSecretKey: jwtSecret,
		InternalKey:  internalKey,
		Encrypter:    encrypter.NewEncrypter(internalKey),
	})

	if err := srv.Run(); err != nil {
		l.Fatalf(ctx, "server.Run: %v", err)
	}
}
