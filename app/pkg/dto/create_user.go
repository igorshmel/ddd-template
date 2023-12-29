package dto

import (
	"github.com/gin-gonic/gin"
)

import validation "github.com/go-ozzo/ozzo-validation/v4"

// CreateUserRequest --
type CreateUserRequest struct {
	UserName string `json:"user_name"`
}

// NewCreateUserRequest is constructor
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

// Parse parses and validates the request
func (ths *CreateUserRequest) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *CreateUserRequest) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.UserName, validation.Required.Error("is required")),
	)
}
