package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/nujabuz/gocourse/internal/infrastructures/database/postgres"
	"github.com/nujabuz/gocourse/pkg/config"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	slog.SetDefault(logger)

	cfg, err := config.Load("config.yaml")
	if err != nil {
		slog.Error("failed to read config.yaml", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbPool, err := postgres.NewPostgres(ctx, cfg.DatabaseConfig)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	res, err := dbPool.Exec(ctx, "INSERT INTO users (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4);",
		"nujabuz",
		"nujabuz123456",
		time.Now(),
		time.Now(),
	)
	if err != nil {
		slog.Error("failed create", "error", err)
	}
	slog.Info("row affected", "row", res.RowsAffected())
}
