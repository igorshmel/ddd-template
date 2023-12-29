package models

import (
	"github.com/lib/pq"
	"time"
)

// ordersTableName - имя таблицы для хранения информации о пользователях
const ordersTableName = "orders"

// OrderModel - модель заказов
type OrderModel struct {
	ID            int64          `gorm:"primaryKey;auto_increment;unique"`         // Базовый инкремент записи
	UUID          string         `gorm:"column:uuid;not null"`                     // Уникальный номер заказа
	UserUUID      string         `gorm:"column:user_uuid"`                         // Уникальный номер пользователя
	TotalPrice    uint64         `gorm:"column:total_price"`                       // Общая цена
	CartItemUUIDs pq.StringArray `gorm:"type:text[]; column:cart_items;not null"`  // Уникальные идентификаторы позиций
	Status        string         `gorm:"column:status;type:order_status;not null"` // Статус заказа
	UpdatedAt     *time.Time     `gorm:"<-:update"`                                // Время обновления записи
	CreatedAt     time.Time      `gorm:"<-:create;index"`                          // Время создания записи
}

// TableName возвращает имя таблицы
func (ths OrderModel) TableName() string {
	return ordersTableName
}

// CreateOrderStatusEnum создает перечисляемый тип для статуса сущности заказа order
func (ths OrderModel) CreateOrderStatusEnum() string {
	query := "DO $$ " +
		"BEGIN " +
		"IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = 'order_status ') THEN " +
		"CREATE TYPE order_status AS ENUM ('new', 'processing', 'completed', 'canceled');" +
		"END IF;" +
		"END$$;"
	return query
}
