package dbo

import (
	"time"
)

// UserDBO - data base object для сущности User --
type UserDBO struct {
	UUID      string
	Name      string
	Balance   int64
	UpdatedAt *time.Time
	CreatedAt time.Time
}
