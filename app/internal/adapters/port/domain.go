package port

import (
	"default_ddd/app/pkg/ddo"
	"default_ddd/app/pkg/logger"
)

//go:generate mockgen -destination=../../mocks/domain.go -package=mocks default_ddd/app/internal/adapters/port CartItemPort,ProductPort,UserPort,OrderPort

// CartItemPort --
type CartItemPort interface {
	ReadCartItem() *ddo.CartItemDDO
	ReadCartItems() []ddo.CartItemDDO
	FillCartItem(*ddo.CartItemDDO)
	FillCartItems([]ddo.CartItemDDO)
	CreateCartItem(logger.Logger, *ddo.CartItemDDO, *ddo.ProductDDO) error
	BatchChangeStatusToCompleted(logger.Logger) error
}

// OrderPort --
type OrderPort interface {
	ReadOrder() *ddo.OrderDDO
	CreateOrder(logger.Logger, *ddo.OrderDDO) error
}

// ProductPort --
type ProductPort interface {
	CreateProduct(*ddo.ProductDDO)
	ReadProduct() *ddo.ProductDDO
}

// UserPort --
type UserPort interface {
	CreateUser(*ddo.UserDDO) *ddo.UserDDO
	FillUser() *ddo.UserDDO
	LoadUser(*ddo.UserDDO)
}
