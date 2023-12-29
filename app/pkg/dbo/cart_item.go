package dbo

import "time"

// CartItemDBO - data base object для сущности CartItem --
type CartItemDBO struct {
	UUID        string
	ProductUUID string
	Quantity    int
	Status      string
	Price       uint64
	UpdatedAt   *time.Time
	CreatedAt   time.Time
}
