package ddo

import (
	status "default_ddd/app/pkg/vars/statuses"
	"time"
)

// OrderDDO --
type OrderDDO struct {
	UUID          string
	UserUUID      string
	CartItemUUIDs []string
	TotalPrice    uint64
	Status        status.CartItemStatus
	UpdatedAt     *time.Time
	CreatedAt     time.Time
}
