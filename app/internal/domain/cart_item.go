package domain

import (
	"default_ddd/app/pkg/ddo"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/logger"
	status "default_ddd/app/pkg/vars/statuses"
	"github.com/google/uuid"
	"time"
)

// CartItem --
type CartItem struct {
	uuid        string                // Уникальный идентификатор позиции
	productUUID string                // Уникальный идентификатор товара
	quantity    int                   // Количество товара в позиции
	status      status.CartItemStatus // Статус позиции
	price       uint64                // Цена заказа
	updatedAt   *time.Time            // Время изменения записи реестра
	createdAt   time.Time             // Время внесения записи в реестр
}

// newCartItem
func (ths *Domain) newCartItem() {
	ths.cartItem.uuid = uuid.New().String()
	ths.cartItem.status = status.CartItemNew
	ths.cartItemCreatedAt()
	ths.cartItemUpdatedAt()
}

// ReadCartItem --
func (ths *Domain) ReadCartItem() *ddo.CartItemDDO {
	return &ddo.CartItemDDO{
		UUID:        ths.cartItem.uuid,
		ProductUUID: ths.cartItem.productUUID,
		Quantity:    ths.cartItem.quantity,
		Status:      ths.cartItem.status,
		Price:       ths.cartItem.price,
		UpdatedAt:   ths.cartItem.updatedAt,
		CreatedAt:   ths.cartItem.createdAt,
	}
}

// ReadCartItems --
func (ths *Domain) ReadCartItems() []ddo.CartItemDDO {

	cartItemsDDO := make([]ddo.CartItemDDO, len(ths.cartItems))

	for i, item := range ths.cartItems {
		cartItemsDDO[i].UUID = item.uuid
		cartItemsDDO[i].ProductUUID = item.productUUID
		cartItemsDDO[i].Price = item.price
		cartItemsDDO[i].Status = item.status
		cartItemsDDO[i].Quantity = item.quantity
		cartItemsDDO[i].CreatedAt = item.createdAt
		cartItemsDDO[i].UpdatedAt = item.updatedAt
	}
	return cartItemsDDO
}

// FillCartItem --
func (ths *Domain) FillCartItem(req *ddo.CartItemDDO) {
	ths.cartItem.uuid = req.UUID
	ths.cartItem.productUUID = req.ProductUUID
	ths.cartItem.quantity = req.Quantity
	ths.cartItem.status = req.Status
	ths.cartItem.price = req.Price
	ths.cartItem.updatedAt = req.UpdatedAt
	ths.cartItem.createdAt = req.CreatedAt
}

// FillCartItems --
func (ths *Domain) FillCartItems(requests []ddo.CartItemDDO) {
	for _, req := range requests {
		ths.cartItems = append(ths.cartItems, CartItem{
			uuid:        req.UUID,
			productUUID: req.ProductUUID,
			quantity:    req.Quantity,
			status:      req.Status,
			price:       req.Price,
			updatedAt:   req.UpdatedAt,
			createdAt:   req.CreatedAt,
		})
	}
}

// getCartItemsTotalPrice --
func (ths *Domain) getCartItemsTotalPrice() uint64 {
	var totalPrice uint64

	for _, item := range ths.cartItems {
		totalPrice = totalPrice + item.price
	}

	return totalPrice
}

// setCartItemsCompleted --
func (ths *Domain) setCartItemsCompleted(l logger.Logger) error {
	log := l.WithMethod("setCartItemsCompleted domain")

	for _, item := range ths.cartItems {
		if item.status != status.CartItemProcessing {
			log.Error("invalid status (status = %s), but must be PROCESSING", item.status)
			return errs.ErrInvalidStatus
		}
		item.status = status.CartItemCompleted
	}

	return nil
}

// CreateCartItem --
func (ths *Domain) CreateCartItem(l logger.Logger, cartItem *ddo.CartItemDDO, product *ddo.ProductDDO) error {
	log := l.WithMethod("CreateCartItem domain")
	ths.newCartItem()
	ths.loadProduct(product)

	if err := ths.reduceTheQuantityProduct(log, product, cartItem.Quantity); err != nil {
		return err
	}

	if ths.cartItem.status != status.CartItemNew {
		log.Error("invalid status (status = %s), but must be NEW", ths.cartItem.status)
		return errs.ErrInvalidStatus
	}

	ths.cartItem.price = uint64(cartItem.Quantity) * ths.product.price
	ths.cartItem.status = status.CartItemProcessing
	ths.cartItem.productUUID = cartItem.ProductUUID
	ths.cartItem.quantity = cartItem.Quantity

	return nil
}

// BatchChangeStatusToCompleted --
func (ths *Domain) BatchChangeStatusToCompleted(l logger.Logger) error {
	log := l.WithMethod("BatchChangeStatusToCompleted domain")

	for _, r := range ths.cartItems {
		if r.status != status.CartItemProcessing {
			log.Error("invalid status (status = %s), but must be processing ", r.status)
			return errs.ErrInvalidStatus
		}
		r.status = status.CartItemCompleted
	}
	return nil
}

// cartItemUpdatedAt --
func (ths *Domain) cartItemUpdatedAt() {
	t := time.Now()
	ths.cartItem.updatedAt = &t
}

// caretItemCreatedAt --
func (ths *Domain) cartItemCreatedAt() {
	t := time.Now()
	ths.cartItem.createdAt = t
}
