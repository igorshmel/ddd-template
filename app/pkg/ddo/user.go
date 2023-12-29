package ddo

import (
	"time"
)

// UserDDO --
type UserDDO struct {
	UUID      string
	Name      string
	Balance   int64
	UpdatedAt *time.Time
	CreatedAt time.Time
}
