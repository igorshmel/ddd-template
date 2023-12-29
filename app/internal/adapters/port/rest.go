package port

import "github.com/gin-gonic/gin"

//go:generate mockgen -destination=../../mocks/rest.go -package=mocks default_ddd/app/internal/adapters/port CreateUserEndpoint,CreateProductEndpoint,CreateCartItemEndpoint,CreateOrderEndpoint

// CreateUserEndpoint --
type CreateUserEndpoint interface {
	ExecuteCreateUserEndpoint(ginContext *gin.Context)
}

// CreateProductEndpoint --
type CreateProductEndpoint interface {
	ExecuteCreateProductEndpoint(ginContext *gin.Context)
}

// CreateCartItemEndpoint --
type CreateCartItemEndpoint interface {
	ExecuteCreateCartItemEndpoint(ginContext *gin.Context)
}

// CreateOrderEndpoint --
type CreateOrderEndpoint interface {
	ExecuteCreateOrderEndpoint(ginContext *gin.Context)
}
