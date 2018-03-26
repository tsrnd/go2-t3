package main

import (
	"fmt"

	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tsrnd/go2-t3/config"
	"github.com/tsrnd/go2-t3/model"
	"github.com/tsrnd/go2-t3/router"
)

func main() {
	fmt.Print("HELLO TEAM 3")

	log.Printf("Server started on: http://localhost%s", os.Getenv("SERVER_PORT"))

	r := router.Route()

	http.ListenAndServe(os.Getenv("SERVER_PORT"), r)
}

func init() {
	config.SetEnv()
	db := config.ConnectDB()
	model.SetDatabase(db)
}
