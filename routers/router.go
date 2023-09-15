package routers

import (
	"github.com/gorilla/mux"
	"github.com/washington-shoji/gobare/handlers"
	"github.com/washington-shoji/gobare/helpers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", helpers.MakeHTTPHandleFunc(handlers.HandleHealthCheck))

	return router
}
