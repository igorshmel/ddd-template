package rest

import (
	"default_ddd/app/cmd/default_ddd/middleware"
	"default_ddd/app/internal/adapters/port"
	"default_ddd/app/pkg/dto"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCartItemEndpoint --
type CreateCartItemEndpoint struct {
	log     logger.Logger
	usecase port.CreateCartItemUseCase
}

// NewCreateCartItemEndpoint --
func NewCreateCartItemEndpoint(usecase port.CreateCartItemUseCase, log logger.Logger) port.CreateCartItemEndpoint {
	return CreateCartItemEndpoint{
		log:     log,
		usecase: usecase,
	}
}

// ExecuteCreateCartItemEndpoint is handler
func (ths CreateCartItemEndpoint) ExecuteCreateCartItemEndpoint(ctx *gin.Context) {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("CreateCartItem endpoint")

	req := dto.NewCreateCartItemRequest()

	// request parse
	if err := req.Parse(ctx); err != nil {
		log.Error("unable to parse a request: %s", err)
		ctx.JSON(http.StatusInternalServerError, errs.ErrParseRequest)
		return
	}

	// validate request
	if err := req.Validate(); err != nil {
		log.Error("error of validation: %s", err)
		ctx.JSON(http.StatusInternalServerError, errs.ErrSyntax)
		return
	}

	// call usecase
	err := ths.usecase.Execute(ctx, req)
	if err != nil {
		log.Error("%s", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
