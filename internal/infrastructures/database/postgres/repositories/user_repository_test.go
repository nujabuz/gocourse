package repositories_test

import (
	"context"
	"testing"

	"github.com/nujabuz/gocourse/internal/domain/entities"
	"github.com/nujabuz/gocourse/internal/infrastructures/database/postgres"
	"github.com/nujabuz/gocourse/internal/infrastructures/database/postgres/repositories"
	"github.com/nujabuz/gocourse/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	cfg, err := config.Load("./../../../../../config.yaml")
	if err != nil {
		panic(err)
	}

	dbPool, err := postgres.NewPostgres(context.Background(), cfg.DatabaseConfig)
	if err != nil {
		panic(err)
	}

	userRepository := repositories.NewUserRepository(dbPool)

	t.Run("create success", func(t *testing.T) {
		user := entities.NewUser("nujabuzzzz", "nujabuz123456789")
		err := user.Validate()

		assert.NoError(t, err)
		assert.Nil(t, err)

		err = userRepository.Create(context.Background(), user)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
