package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//ConnectDB config connect db
func ConnectDB() *gorm.DB {
	dbConnect := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	db, err := gorm.Open(dbConnect, "host="+dbHost+" port="+dbPort+" user="+dbUser+" dbname="+dbName+" password="+dbPass)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.LogMode(true)
	return db
}
