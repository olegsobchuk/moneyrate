package router

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/olegsobchuk/moneyrate/controllers"
)

// MakeRouter creates routes
func MakeRouter() http.Handler {
	dir := "./public"
	router := mux.NewRouter()
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(dir))))
	router.HandleFunc("/", controller.ExchangeIndex).Methods("GET")
	router.HandleFunc("/", controller.ExchangeFindRate).Methods("POST")

	return router
}
