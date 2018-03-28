package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tsrnd/go2-t3/handler"
)

func Route() *mux.Router {
	var bh handler.BlogHandler
	r := mux.NewRouter()

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.HandleFunc("/", bh.Index).Methods("GET")
	r.HandleFunc("/create", bh.Create).Methods("GET")
	r.HandleFunc("/store", bh.Store).Methods("POST")

	return r
}
