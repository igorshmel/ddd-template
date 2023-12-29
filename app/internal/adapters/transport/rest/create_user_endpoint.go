package rest

import (
	"net/http"

	"default_ddd/app/cmd/default_ddd/middleware"
	"default_ddd/app/internal/adapters/port"
	"default_ddd/app/pkg/dto"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/logger"
	"github.com/gin-gonic/gin"
)

// CreateUserEndpoint --
type CreateUserEndpoint struct {
	log     logger.Logger
	usecase port.CreateUserUseCase
}

// NewCreateUserEndpoint --
func NewCreateUserEndpoint(usecase port.CreateUserUseCase, log logger.Logger) port.CreateUserEndpoint {
	return CreateUserEndpoint{
		log:     log,
		usecase: usecase,
	}
}

// ExecuteCreateUserEndpoint is handler
func (ths CreateUserEndpoint) ExecuteCreateUserEndpoint(ctx *gin.Context) {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("CreateUser endpoint")

	req := dto.NewCreateUserRequest()

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
