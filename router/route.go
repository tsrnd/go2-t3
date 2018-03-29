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

	r.HandleFunc("/{id}/edit", bh.Edit).Methods("GET")
	r.HandleFunc("/{id}/update", bh.Update).Methods("POST")
	r.HandleFunc("/blogs", bh.Index).Methods("GET")

	return r
}
