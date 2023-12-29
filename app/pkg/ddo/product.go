package ddo

import (
	"time"
)

// ProductDDO --
type ProductDDO struct {
	UUID      string
	Title     string
	Price     uint64
	Quantity  int
	UpdatedAt *time.Time
	CreatedAt time.Time
}
