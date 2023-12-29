package dto

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CreateCartItemRequest --
type CreateCartItemRequest struct {
	ProductUUID string `json:"product_uuid"`
	Quantity    int    `json:"quantity"`
}

// NewCreateCartItemRequest is constructor
func NewCreateCartItemRequest() *CreateCartItemRequest {
	return &CreateCartItemRequest{}
}

// Parse parses and validates the request
func (ths *CreateCartItemRequest) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *CreateCartItemRequest) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.ProductUUID, validation.Required.Error("is required")),
		validation.Field(&ths.Quantity, validation.Required.Error("is required")),
	)
}
