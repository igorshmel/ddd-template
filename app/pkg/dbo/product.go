package dbo

import "time"

// ProductDBO - data base object для сущности Product --
type ProductDBO struct {
	UUID      string
	Title     string
	Price     uint64
	Quantity  int
	UpdatedAt *time.Time
	CreatedAt time.Time
}
