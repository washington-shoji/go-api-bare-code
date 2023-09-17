package routers

import (
	"github.com/gorilla/mux"
	"github.com/washington-shoji/gobare/handlers"
	"github.com/washington-shoji/gobare/helpers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", helpers.MakeHTTPHandleFunc(handlers.HandleHealthCheck))
	router.HandleFunc("/account", helpers.MakeHTTPHandleFunc(handlers.HandleAccount))
	router.HandleFunc("/account/{id}", helpers.MakeHTTPHandleFunc(handlers.HandleAccountByID))

	return router
}
