package config_test

import (
	"os"
	"testing"

	"github.com/nujabuz/gocourse/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("success load", func(t *testing.T) {
		cfg, err := config.Load("./../../config.yaml")

		assert.NoError(t, err)
		assert.Nil(t, err)

		assert.NotEqual(t, "", cfg.DatabaseConfig.ConnectionString)
	})

	t.Run("file not found", func(t *testing.T) {
		cfg, err := config.Load("./")

		pathError := new(os.PathError)

		assert.Error(t, err)
		assert.ErrorAs(t, err, &pathError)

		assert.Nil(t, cfg)
	})
}
