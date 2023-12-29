package domain

import (
	"default_ddd/app/pkg/ddo"
	"default_ddd/app/pkg/logger"
	status "default_ddd/app/pkg/vars/statuses"
	"github.com/google/uuid"
	"time"
)

// Order --
type Order struct {
	uuid          string             // Уникальный идентификатор заказа
	userUUID      string             // Уникальный идентификатор пользователя
	cartItemUUIDs []string           // Массив идентификаторов позиций заказа
	totalPrice    uint64             // Общая цена позиций / Цена заказа
	status        status.OrderStatus // Статус заказа
	updatedAt     *time.Time         // Время изменения записи реестра
	createdAt     time.Time          // Время внесения записи в реестр
}

// newOrder
func (ths *Domain) newOrder() {
	ths.order.uuid = uuid.New().String()
	ths.order.status = status.OrderNew
	ths.orderCreatedAt()
	ths.orderUpdatedAt()
}

// ReadOrder --
func (ths *Domain) ReadOrder() *ddo.OrderDDO {
	return &ddo.OrderDDO{
		UUID:          ths.order.uuid,
		UserUUID:      ths.order.userUUID,
		TotalPrice:    ths.order.totalPrice,
		CartItemUUIDs: ths.order.cartItemUUIDs,
		Status:        status.CartItemStatus(ths.order.status.String()),
		UpdatedAt:     ths.order.updatedAt,
		CreatedAt:     ths.order.createdAt,
	}
}

// CreateOrder --
func (ths *Domain) CreateOrder(l logger.Logger, req *ddo.OrderDDO) error {
	log := l.WithMethod("CreateOrder domain")
	ths.newOrder()

	if err := ths.setCartItemsCompleted(log); err != nil {
		return err
	}

	totalPrice := ths.getCartItemsTotalPrice()
	if err := ths.reduceUserBalance(log, totalPrice); err != nil {
		return err
	}

	ths.order.userUUID = req.UserUUID
	ths.order.totalPrice = totalPrice
	ths.order.cartItemUUIDs = req.CartItemUUIDs

	return nil
}

// orderUpdatedAt --
func (ths *Domain) orderUpdatedAt() {
	t := time.Now()
	ths.order.updatedAt = &t
}

// orderCreatedAt --
func (ths *Domain) orderCreatedAt() {
	t := time.Now()
	ths.order.createdAt = t
}
