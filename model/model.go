package model

import (
	"github.com/jinzhu/gorm"
)

//DBCon dbcon
var DBCon *gorm.DB

//SetDatabase return DBCon
func SetDatabase(database *gorm.DB) {
	DBCon = database
}
