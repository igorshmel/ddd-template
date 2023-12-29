package routes

import (
	"default_ddd/app/internal/adapters/repository"
	"default_ddd/app/internal/adapters/transport/rest"
	"default_ddd/app/internal/domain"
	"default_ddd/app/internal/usecase/api"
	"default_ddd/app/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func registerRoutes(dom *domain.Ports, log logger.Logger, g *gin.Engine, repo *repository.Repository) {
	apiGroup := g.Group("/api")
	v1 := apiGroup.Group("/v1")

	// Создание usecase
	createUserUseCase := api.NewCreateUserUseCase(log, repo.GetPersister(), repo.GetExtractor(), dom.GetUserPort())
	createProductUseCase := api.NewCreateProductUseCase(log, repo.GetPersister(), repo.GetExtractor(), dom.GetProductPort())
	createCartItemUseCase := api.NewCreateCartItemUseCase(log, repo.GetPersister(), repo.GetExtractor(), dom.GetCartItemPort(), dom.GetProductPort())
	createOrderUseCase := api.NewCreateOrderUseCase(log, dom.GetOrderPort(), dom.GetUserPort(), dom.GetCartItemPort(), repo.GetExtractor(), repo.GetPersister())

	// Создание обработчиков запросов
	createUserEndpoint := rest.NewCreateUserEndpoint(createUserUseCase, log)
	createProductEndpoint := rest.NewCreateProductEndpoint(createProductUseCase, log)
	createCartItemEndpoint := rest.NewCreateCartItemEndpoint(createCartItemUseCase, log)
	createOrderEndpoint := rest.NewCreateOrderEndpoint(createOrderUseCase, log)

	// Регистрация обработчиков запросов
	v1.POST("/user/", createUserEndpoint.ExecuteCreateUserEndpoint)
	v1.POST("/product/", createProductEndpoint.ExecuteCreateProductEndpoint)
	v1.POST("/cart_item/", createCartItemEndpoint.ExecuteCreateCartItemEndpoint)
	v1.POST("/order/", createOrderEndpoint.ExecuteCreateOrderEndpoint)
}

// Module ..
var Module = fx.Options(fx.Invoke(registerRoutes))
