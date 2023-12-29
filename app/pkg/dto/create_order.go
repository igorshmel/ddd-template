package dto

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CreateOrderRequest --
type CreateOrderRequest struct {
	UserUUID      string   `json:"user_uuid"`
	CartItemUUIDs []string `json:"cart_item_uuids"`
}

// NewCreateOrderRequest is constructor
func NewCreateOrderRequest() *CreateOrderRequest {
	return &CreateOrderRequest{}
}

// Parse parses and validates the request
func (ths *CreateOrderRequest) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *CreateOrderRequest) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.UserUUID, validation.Required.Error("is required")),
		validation.Field(&ths.CartItemUUIDs, validation.Required.Error("is required")),
	)
}
