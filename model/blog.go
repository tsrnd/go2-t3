package model

import (
	"os"
	"time"
)

// Blog struct
type Blog struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	Description string
	Detail      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// TableName name
func (Blog) TableName() string {
	return os.Getenv("DB_DATABASE") + ".blogs"
}
