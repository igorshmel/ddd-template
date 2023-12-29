package models

import "time"

// cartItemTableName - имя таблицы для хранения информации о корзине покупателя
const cartItemTableName = "cart_items"

// CartItemModel - модель корзины
type CartItemModel struct {
	ID          int64      `gorm:"primaryKey;auto_increment;unique"`             // Базовый инкремент записи
	UUID        string     `gorm:"column:uuid;not null"`                         // Уникальный номер заказа
	ProductUUID string     `gorm:"column:product_uuid"`                          // Уникальный номер продукта
	Quantity    int        `gorm:"column:quantity"`                              // Количество продукта в позиции
	Status      string     `gorm:"column:status;type:cart_item_status;not null"` // Статус позиции в корзине заказов
	Price       uint64     `gorm:"column:price"`                                 // Цена заказа
	UpdatedAt   *time.Time `gorm:"<-:update"`                                    // Время обновления записи
	CreatedAt   time.Time  `gorm:"<-:create;index"`                              // Время создания записи
}

// TableName возвращает имя таблицы
func (ths CartItemModel) TableName() string {
	return cartItemTableName
}

// CreateCartItemStatusEnum создает перечисляемый тип для статуса сущности позиции cart_items
func (ths CartItemModel) CreateCartItemStatusEnum() string {
	query := "DO $$ " +
		"BEGIN " +
		"IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = 'cart_item_status ') THEN " +
		"CREATE TYPE cart_item_status AS ENUM ('new', 'processing', 'completed', 'canceled');" +
		"END IF;" +
		"END$$;"
	return query
}
