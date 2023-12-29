package domain

import (
	"default_ddd/app/pkg/ddo"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/logger"
	"github.com/google/uuid"
	"time"
)

// Product --
type Product struct {
	uuid      string     // Уникальный идентификатор товара
	title     string     // Наименование товара
	price     uint64     // Цена товара
	quantity  int        // Количество товара на складе
	updatedAt *time.Time // Время изменения записи реестра
	createdAt time.Time  // Время внесения записи в реестр
}

// newProduct --
func (ths *Domain) newProduct() {
	ths.product.uuid = uuid.New().String()
	ths.productCreatedAt()
	ths.productUpdatedAt()
}

// ReadProduct --
func (ths *Domain) ReadProduct() *ddo.ProductDDO {
	return &ddo.ProductDDO{
		UUID:      ths.product.uuid,
		Title:     ths.product.title,
		Price:     ths.product.price,
		Quantity:  ths.product.quantity,
		UpdatedAt: ths.product.updatedAt,
		CreatedAt: ths.product.createdAt,
	}
}

// loadProduct --
func (ths *Domain) loadProduct(req *ddo.ProductDDO) {
	ths.product.uuid = req.UUID
	ths.product.title = req.Title
	ths.product.price = req.Price
	ths.product.quantity = req.Quantity
	ths.product.updatedAt = req.UpdatedAt
	ths.product.createdAt = req.CreatedAt
}

// CreateProduct --
func (ths *Domain) CreateProduct(ddo *ddo.ProductDDO) {
	ths.newProduct()
	ths.product.title = ddo.Title
	ths.product.quantity = ddo.Quantity
	ths.product.price = ddo.Price
}

// UpdateProduct --
func (ths *Domain) reduceTheQuantityProduct(l logger.Logger, ddo *ddo.ProductDDO, quantity int) error {
	log := l.WithMethod("UpdateProduct domain")

	if ddo.Quantity < quantity {
		log.Error("the quantity of product is not enough (productQuantity = %d), (cartItemQuantity = %d) ", ddo.Quantity, quantity)
		return errs.ErrQuantityIsNotEnough
	}

	ths.product.quantity = ddo.Quantity - quantity
	ths.productUpdatedAt()

	return nil
}

// productUpdatedAt --
func (ths *Domain) productUpdatedAt() {
	t := time.Now()
	ths.product.updatedAt = &t
}

// productCreatedAt --
func (ths *Domain) productCreatedAt() {
	t := time.Now()
	ths.product.createdAt = t
}
