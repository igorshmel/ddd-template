package dto

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CreateProductRequest --
type CreateProductRequest struct {
	Title    string `json:"title"`
	Price    uint64 `json:"price"`
	Quantity int    `json:"quantity"`
}

// NewCreateProductRequest is constructor
func NewCreateProductRequest() *CreateProductRequest {
	return &CreateProductRequest{}
}

// Parse parses and validates the request
func (ths *CreateProductRequest) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *CreateProductRequest) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.Title, validation.Required.Error("is required")),
		validation.Field(&ths.Price, validation.Required.Error("is required")),
		validation.Field(&ths.Quantity, validation.Required.Error("is required")),
	)
}
