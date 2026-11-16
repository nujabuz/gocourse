// Package postgres
package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nujabuz/gocourse/pkg/config"
)

func NewPostgres(ctx context.Context, dbCfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbCfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = int32(dbCfg.MaxConnections)
	cfg.MaxConnIdleTime = time.Second * time.Duration(dbCfg.MaxConnectionsIdleTime)
	cfg.MaxConnLifetime = time.Second * time.Duration(dbCfg.MaxConnectionsLifetime)
	cfg.MinIdleConns = int32(dbCfg.MinIdleConnections)

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
