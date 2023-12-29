package log

import (
	"default_ddd/app/pkg/logger"
	"go.uber.org/fx"
)

// newLogger создаёт новый экземпляр Logger
func newLogger() logger.Logger {
	return logger.New(true)
}

// Module ..
var Module = fx.Options(fx.Provide(newLogger))
