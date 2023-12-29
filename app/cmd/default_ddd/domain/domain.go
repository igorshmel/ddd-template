package domain

import (
	"default_ddd/app/internal/domain"
	"default_ddd/app/pkg/logger"
	"go.uber.org/fx"
)

func newDom(log logger.Logger) (*domain.Ports, error) {
	dom, err := domain.NewPorts(
		domain.WithDefaultDomain(log))
	if err != nil {
		log.Fatal("failed initialize domain with error: %s", err.Error())
	}
	return dom, err
}

// Module ..
var Module = fx.Options(fx.Provide(newDom))
