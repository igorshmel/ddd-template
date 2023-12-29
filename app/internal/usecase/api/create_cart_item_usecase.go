package api

import (
	"context"
	"default_ddd/app/pkg/dbo"
	"default_ddd/app/pkg/dto"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/mapping"
	"errors"

	"default_ddd/app/cmd/default_ddd/middleware"
	"default_ddd/app/internal/adapters/port"
	"default_ddd/app/pkg/logger"
)

// CreateCartItemUseCase --
type CreateCartItemUseCase struct {
	log       logger.Logger
	persister port.Persister
	extractor port.Extractor
	cartItem  port.CartItemPort
	product   port.ProductPort
}

// NewCreateCartItemUseCase --
func NewCreateCartItemUseCase(log logger.Logger, persister port.Persister, extractor port.Extractor, cartItem port.CartItemPort, product port.ProductPort) port.CreateCartItemUseCase {
	return CreateCartItemUseCase{log: log, persister: persister, extractor: extractor, cartItem: cartItem, product: product}
}

// Execute --
func (ths CreateCartItemUseCase) Execute(ctx context.Context, cartItemDTO *dto.CreateCartItemRequest) error {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("CreateCartItem usecase")

	// -- Чтение данных -- загружает состояние доменных моделей
	// ---------------------------------------------------------------------------------------------------------------------------
	productDBO, err := ths.extractor.GetProductByUUID(&dbo.ProductDBO{UUID: cartItemDTO.ProductUUID})
	if errors.Is(err, errs.ErrNotFound) {
		log.Error("product with (productUUID: %s) not found", err.Error())
		return errs.ErrNotFound
	}
	if err != nil {
		log.Error("failed to get product (productUUID: %s) with error: %s", cartItemDTO.ProductUUID, err.Error())
		return errs.ErrGetProductFailed
	}

	// -- Бизнес логика -- создание позиции заказа на уровне домена
	// ---------------------------------------------------------------------------------------------------------------------------
	cartItemDDO := mapping.CartItemDTOtoDDO(cartItemDTO)
	productDDO := mapping.ProductDBOtoDDO(productDBO)

	if err = ths.cartItem.CreateCartItem(log, cartItemDDO, productDDO); err != nil {
		return err
	}

	cartItemDDO = ths.cartItem.ReadCartItem()
	productDDO = ths.product.ReadProduct()

	// -- Инфраструктурная логика -- запись состояний доменных моделей в БД
	// ---------------------------------------------------------------------------------------------------------------------------
	cartItemDBO := mapping.CartItemDDOtoDBO(cartItemDDO)
	productDBO = mapping.ProductDDOtoDBO(productDDO)

	if err = ths.persister.UnitOfWork(func(tx port.Persister) error { // единица работы, транзакция БД

		if err = ths.persister.UpdateProduct(productDBO, cartItemDBO); err != nil {
			log.Error("failed to update product with error: %s", err.Error())
			return errs.ErrUpdateProduct
		}

		if err = ths.persister.CreateCartItem(cartItemDBO); err != nil {
			log.Error("failed to create cartItem (UUID: %s) with error: %s", cartItemDBO.UUID, err.Error())
			return errs.ErrCreateCartItemFailed
		}

		return nil
	}); err != nil {
		return err
	}

	log.Debug("create an cartItem (uuid: %s)", cartItemDDO.UUID)
	return nil
}
