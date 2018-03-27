package config

import (
	"os"

	"github.com/jinzhu/gorm"
)

//ConnectDB config connect db
func ConnectDB() *gorm.DB {
	dbConnect := os.Getenv("DB_CONNECTION")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	db, err := gorm.Open(dbConnect, dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
