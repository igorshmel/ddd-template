package bundlefx

import (
	"context"
	"fmt"

	"default_ddd/app/cmd/default_ddd/configs"
	"default_ddd/app/cmd/default_ddd/domain"
	"default_ddd/app/cmd/default_ddd/ginserver"
	"default_ddd/app/cmd/default_ddd/log"
	"default_ddd/app/cmd/default_ddd/middleware"
	"default_ddd/app/cmd/default_ddd/repo"
	"default_ddd/app/cmd/default_ddd/routes"
	"default_ddd/app/pkg/config"

	"default_ddd/app/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module собирает все зависимости для запуска приложения
var Module = fx.Options(
	ginserver.Module,
	log.Module,
	configs.Module,
	domain.Module,
	routes.Module,
	repo.Module,

	fx.Invoke(manageServer),
	fx.Invoke(setGinMiddlewares),
	fx.Invoke(setGinLogger),
)

// manageServer управляет запуском и остановкой сервера
func manageServer(lc fx.Lifecycle, log logger.Logger, g *gin.Engine, cfg config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			port := cfg.App.Port
			log.Info("Starting application in :%s", port)

			go func() {
				err := g.Run(":" + port)
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			log.Info("Stopping application")
			return nil
		},
	})
}

// setGinLogger sets standard logger
func setGinLogger(router *gin.Engine) {
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		reqID := param.Request.Header.Get(middleware.RequestIDKey)
		if len(reqID) == 0 {
			reqID = "empty request id"
		}

		return fmt.Sprintf("%s [GIN]  [%v] |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			reqID,
			param.StatusCodeColor(), param.StatusCode, param.ResetColor(),
			param.Latency,
			param.ClientIP,
			param.MethodColor(), param.Method, param.ResetColor(),
			param.Path,
			param.ErrorMessage,
		)
	}))
}

// setGinMiddlewares sets middlewares
func setGinMiddlewares(router *gin.Engine) {
	// middleware CORS settings
	router.Use(cors.Default())
	// middleware with request id for each request
	router.Use(middleware.MakeRequestIDGinMiddleware())
	// recovery middleware
	router.Use(gin.Recovery())
}
