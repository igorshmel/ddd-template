package ginserver

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func newGin() *gin.Engine {
	router := gin.Default()
	return router
}

// Module --
var Module = fx.Options(fx.Provide(newGin))
