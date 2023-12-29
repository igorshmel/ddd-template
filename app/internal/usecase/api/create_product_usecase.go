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

// CreateProductUseCase --
type CreateProductUseCase struct {
	log       logger.Logger
	persister port.Persister
	extractor port.Extractor
	product   port.ProductPort
}

// NewCreateProductUseCase --
func NewCreateProductUseCase(log logger.Logger, persister port.Persister, extractor port.Extractor, product port.ProductPort) port.CreateProductUseCase {
	return CreateProductUseCase{log: log, persister: persister, extractor: extractor, product: product}
}

// Execute --
func (ths CreateProductUseCase) Execute(ctx context.Context, productDTO *dto.CreateProductRequest) error {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("CreateProduct usecase")

	// -- Бизнес логика -- создание продукта на уровне домена
	// ---------------------------------------------------------------------------------------------------------------------------
	ths.product.CreateProduct(
		mapping.ProductDTOtoDDO(productDTO),
	)
	productDDO := ths.product.ReadProduct()

	// -- Инфраструктурная логика -- запись состояний доменных моделей в БД
	// ---------------------------------------------------------------------------------------------------------------------------
	productDBO := mapping.ProductDDOtoDBO(productDDO)
	if err := ths.persister.CreateProduct(productDBO); err != nil {
		log.Error("failed to create product (UUID: %s) with error: %s", productDBO.UUID, err.Error())
		return errs.ErrCreateProductFailed
	}

	log.Debug("create an product (uuid: %s)", productDDO.UUID)
	return nil
}
