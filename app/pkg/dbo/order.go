package dbo

import (
	status "default_ddd/app/pkg/vars/statuses"
	"time"
)

// OrderDBO --
type OrderDBO struct {
	UUID          string             // Уникальный идентификатор заказа
	UserUUID      string             // Уникальный идентификатор пользователя
	CartItemUUIDs []string           // Массив идентификаторов позиций заказа
	TotalPrice    uint64             // Общая цена позиций / Цена заказа
	Status        status.OrderStatus // Статус заказа
	UpdatedAt     *time.Time         // Время изменения записи реестра
	CreatedAt     time.Time          // Время внесения записи в реестр
}
