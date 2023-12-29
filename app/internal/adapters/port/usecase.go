package port

import (
	"context"

	"default_ddd/app/pkg/dto"
)

//go:generate mockgen -destination=../../mocks/usecase.go -package=mocks default_ddd/app/internal/adapters/port CreateUserUseCase,CreateProductUseCase,CreateCartItemUseCase,CreateOrderUseCase

// CreateUserUseCase --
type CreateUserUseCase interface {
	Execute(context.Context, *dto.CreateUserRequest) error
}

// CreateProductUseCase --
type CreateProductUseCase interface {
	Execute(context.Context, *dto.CreateProductRequest) error
}

// CreateCartItemUseCase --
type CreateCartItemUseCase interface {
	Execute(context.Context, *dto.CreateCartItemRequest) error
}

// CreateOrderUseCase --
type CreateOrderUseCase interface {
	Execute(context.Context, *dto.CreateOrderRequest) error
}
