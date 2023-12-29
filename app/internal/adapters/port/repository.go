package port

import "default_ddd/app/pkg/dbo"

//go:generate mockgen -destination=../../mocks/repository.go -package=mocks default_ddd/app/internal/adapters/port Extractor,Persister

// Extractor - объект для извлечения данных из БД
type Extractor interface {
	GetUserByUUID(string) (*dbo.UserDBO, error)
	GetProductByUUID(*dbo.ProductDBO) (*dbo.ProductDBO, error)
	GetCartItemsByUUIDs([]string) ([]dbo.CartItemDBO, error)
}

// Persister - объект для сохранения данных в БД
type Persister interface {
	CreateUser(*dbo.UserDBO) error
	UpdateUser(*dbo.UserDBO) error
	CreateProduct(*dbo.ProductDBO) error
	UpdateProduct(*dbo.ProductDBO, *dbo.CartItemDBO) error
	CreateCartItem(*dbo.CartItemDBO) error
	UpdateCartItems(*dbo.OrderDBO) error
	CreateOrder(*dbo.OrderDBO) error
	UnitOfWork(func(Persister) error) (err error)
}
