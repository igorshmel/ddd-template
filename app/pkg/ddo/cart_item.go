package ddo

import (
	status "default_ddd/app/pkg/vars/statuses"
	"time"
)

// CartItemDDO --
type CartItemDDO struct {
	UUID        string
	ProductUUID string
	Quantity    int
	Status      status.CartItemStatus
	Price       uint64
	UpdatedAt   *time.Time
	CreatedAt   time.Time
}
