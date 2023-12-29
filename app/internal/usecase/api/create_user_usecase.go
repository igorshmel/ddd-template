package api

import (
	"context"
	"default_ddd/app/cmd/default_ddd/middleware"
	"default_ddd/app/internal/adapters/port"
	"default_ddd/app/pkg/dto"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/logger"
	"default_ddd/app/pkg/mapping"
)

// CreateUserUseCase --
type CreateUserUseCase struct {
	log       logger.Logger
	persister port.Persister
	extractor port.Extractor
	user      port.UserPort
}

// NewCreateUserUseCase --
func NewCreateUserUseCase(log logger.Logger, persister port.Persister, extractor port.Extractor, user port.UserPort) port.CreateUserUseCase {
	return CreateUserUseCase{log: log, persister: persister, extractor: extractor, user: user}
}

// Execute --
func (ths CreateUserUseCase) Execute(ctx context.Context, userDTO *dto.CreateUserRequest) error {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("CreateUser UseCase ")

	// -- Бизнес логика -- создание пользователя на уровне домена
	// ---------------------------------------------------------------------------------------------------------------------------
	userDDO := ths.user.CreateUser(
		mapping.UserDTOtoDDO(userDTO),
	)

	// -- Инфраструктурная логика -- запись состояний доменных моделей в БД
	// ---------------------------------------------------------------------------------------------------------------------------
	userDBO := mapping.UserDDOtoDBO(userDDO)
	if err := ths.persister.CreateUser(userDBO); err != nil {
		log.Error("failed to create user (UUID: %s) with error: %s", userDBO.UUID, err.Error())
		return errs.ErrCreateUserFailed
	}

	log.Debug("create an user (uuid: %s)", userDDO.UUID)

	return nil
}
