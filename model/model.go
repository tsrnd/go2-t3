package model

import "database/sql"

//DBCon dbcon
var DBCon *sql.DB

//SetDatabase return DBCon
func SetDatabase(database *sql.DB) {
	DBCon = database
}
