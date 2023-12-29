package models

import "time"

// productsTableName - имя таблицы для хранения информации о продукте
const productsTableName = "products"

// ProductModel - модель продуктов
type ProductModel struct {
	ID        int64      `gorm:"primaryKey;auto_increment;unique"` // Базовый инкремент записи
	UUID      string     `gorm:"column:uuid;not null"`             // Уникальный номер продукта
	Title     string     `gorm:"column:name"`                      // Наименование продукта
	Price     uint64     `gorm:"column:price"`                     // Цена продукта
	Quantity  int        `gorm:"column:quantity"`                  // Количество на складе
	UpdatedAt *time.Time `gorm:"<-:update"`                        // Время обновления записи
	CreatedAt time.Time  `gorm:"<-:create;index"`                  // Время создания записи
}

// TableName возвращает имя таблицы
func (ths ProductModel) TableName() string {
	return productsTableName
}
