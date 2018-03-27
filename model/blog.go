package model

import (
	"time"
)

// Blog model
type Blog struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	Description string
	Detail      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
