package repo

import (
	"default_ddd/app/internal/adapters/repository"
	"default_ddd/app/pkg/config"
	"default_ddd/app/pkg/logger"
	"go.uber.org/fx"
)

func newRepository(log logger.Logger, cfg config.Config) (*repository.Repository, error) {
	rep, err := repository.NewRepository(
		repository.WithPostgresRepository(cfg, log))
	if err != nil {
		log.Fatal("failed initialize repository with error: %s", err.Error())
	}
	return rep, err
}

// Module ..
var Module = fx.Options(fx.Provide(newRepository))
