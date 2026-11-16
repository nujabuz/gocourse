package postgres_test

import (
	"context"
	"testing"

	"github.com/nujabuz/gocourse/internal/infrastructures/database/postgres"
	"github.com/nujabuz/gocourse/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
	cfg, err := config.Load("./../../../../config.yaml")
	if err != nil {
		panic(err)
	}

	t.Run("success connection", func(t *testing.T) {
		pool, err := postgres.NewPostgres(context.Background(), cfg.DatabaseConfig)

		assert.NoError(t, err)
		assert.Nil(t, err)

		assert.NotNil(t, pool)
	})

	t.Run("invalid connection string", func(t *testing.T) {
		cfg.DatabaseConfig.ConnectionString = "invalid connection string"
		pool, err := postgres.NewPostgres(context.Background(), cfg.DatabaseConfig)

		assert.Error(t, err)
		assert.Nil(t, pool)
	})
}
