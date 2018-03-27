package config

import (
	"database/sql"
	"os"
)

//ConnectDB config connect db
func ConnectDB() *sql.DB {
	dbConnect := os.Getenv("DB_CONNECTION")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	db, err := sql.Open(dbConnect, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	return db
}
