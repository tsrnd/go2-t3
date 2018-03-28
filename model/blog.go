package model

import (
	"os"

	"github.com/jinzhu/gorm"
)

// Blog model
type Blog struct {
	gorm.Model
	Title       string
	Description string
	Detail      string
}

func (Blog) TableName() string {
	return os.Getenv("DB_DATABASE") + ".blogs"
}
