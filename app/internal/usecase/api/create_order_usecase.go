package api

import (
	"context"
	"default_ddd/app/internal/adapters/port"
	"default_ddd/app/pkg/dto"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/mapping"
	"errors"

	"default_ddd/app/cmd/default_ddd/middleware"
	"default_ddd/app/pkg/logger"
)

// CreateOrderUseCase --
type CreateOrderUseCase struct {
	log       logger.Logger
	order     port.OrderPort
	user      port.UserPort
	cartItem  port.CartItemPort
	extractor port.Extractor
	persister port.Persister
}

// NewCreateOrderUseCase --
func NewCreateOrderUseCase(
	log logger.Logger,
	order port.OrderPort,
	user port.UserPort,
	cartItem port.CartItemPort,
	extractor port.Extractor,
	persister port.Persister,
) port.CreateOrderUseCase {
	return CreateOrderUseCase{
		log:       log,
		order:     order,
		user:      user,
		cartItem:  cartItem,
		extractor: extractor,
		persister: persister,
	}
}

// Execute --
func (ths CreateOrderUseCase) Execute(ctx context.Context, req *dto.CreateOrderRequest) error {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("CreateOrder usecase")

	// -- Чтение данных -- загружает состояние доменных моделей
	// ---------------------------------------------------------------------------------------------------------------------------
	cartItemsDBO, err := ths.extractor.GetCartItemsByUUIDs(req.CartItemUUIDs)
	if errors.Is(err, errs.ErrNotFound) {
		log.Error("cartItems not found", err.Error())
		return errs.ErrCartItemsNotFound
	}
	if err != nil {
		log.Error("failed to get array cartItems with error: %s", err.Error())
	}

	log.Info("cart items with total sum: %v", cartItemsDBO)

	userDBO, err := ths.extractor.GetUserByUUID(req.UserUUID)
	if errors.Is(err, errs.ErrNotFound) {
		log.Error("user not found", err.Error())
		return errs.ErrUserNotFound
	}
	if err != nil {
		log.Error("failed to get user with error: %s", err.Error())
	}

	// -- Бизнес логика -- создание заказа на уровне домена
	// ---------------------------------------------------------------------------------------------------------------------------
	userDDO := mapping.UserDBOtoDDO(userDBO)
	ths.user.LoadUser(userDDO)

	cartItemsDDO := mapping.CartItemsDBOtoDDO(cartItemsDBO)
	ths.cartItem.FillCartItems(cartItemsDDO)

	orderDDO := mapping.OrderDTOtoDDO(req)
	err = ths.order.CreateOrder(log, orderDDO)
	if err != nil {
		return err
	}

	userDDO = ths.user.FillUser()
	orderDDO = ths.order.ReadOrder()

	// -- Инфраструктурная логика -- запись состояний доменных моделей в БД
	// ---------------------------------------------------------------------------------------------------------------------------
	if err = ths.persister.UnitOfWork(func(tx port.Persister) error { // единица работы, транзакция БД

		orderDBO := mapping.OrderDDOtoDBO(orderDDO)
		if err = ths.persister.CreateOrder(orderDBO); err != nil {
			log.Error("failed to create order (UUID: %s) with error: %s", orderDBO.UUID, err.Error())
			return errs.ErrCreateOrderFailed
		}

		//cartItemsDBO = mapping.CartItemsDOOtoDBO(cartItemsDDO)
		if err = ths.persister.UpdateCartItems(orderDBO); err != nil {
			log.Error("failed update cartItems with error: %s", err.Error())
			return errs.ErrUpdateCartItems
		}

		userDBO = mapping.UserDDOtoDBO(userDDO)
		if err = ths.persister.UpdateUser(userDBO); err != nil {
			log.Error("failed update user UUID: %s) with error: %s", userDBO.UUID, err.Error())
			return errs.ErrUpdateUser
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
