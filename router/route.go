package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tsrnd/go2-t3/controller"
)

func Route() *mux.Router {
	var bc controller.BlogController
	r := mux.NewRouter()

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.HandleFunc("/", bc.Index).Methods("GET")

	return r
}
