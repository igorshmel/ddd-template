package models

import "time"

// usersTableName - имя таблицы для хранения информации о пользователях
const usersTableName = "users"

// UserModel - модель БД для сущности пользователя
type UserModel struct {
	ID        int64      `gorm:"primaryKey;auto_increment;unique"`      // Базовый инкремент записи
	UUID      string     `gorm:"column:uuid;type:uuid;unique;not null"` // Уникальный номер пользователя
	Name      string     `gorm:"column:name"`                           // Наименование пользователя
	Balance   int64      `gorm:"column:balance"`                        // Баланс денег на счете пользователя
	UpdatedAt *time.Time `gorm:"<-:update"`                             // Время обновления записи
	CreatedAt time.Time  `gorm:"<-:create;index"`                       // Время создания записи
}

// TableName возвращает имя таблицы
func (ths UserModel) TableName() string {
	return usersTableName
}
