package routers

import (
	"github.com/gorilla/mux"
	"github.com/washington-shoji/gobare/handlers"
	"github.com/washington-shoji/gobare/helpers"
	"github.com/washington-shoji/gobare/middleware"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", helpers.MakeHTTPHandleFunc(handlers.HandleLogin))
	router.HandleFunc("/health-check", helpers.MakeHTTPHandleFunc(handlers.HandleHealthCheck))
	router.HandleFunc("/account", helpers.MakeHTTPHandleFunc(handlers.HandleAccount))
	router.HandleFunc("/account/{id}", middleware.JWTAuth(helpers.MakeHTTPHandleFunc(handlers.HandleAccountByID)))
	router.HandleFunc("/author", helpers.MakeHTTPHandleFunc(handlers.HandleAuthor))
	router.HandleFunc("/author/{id}", helpers.MakeHTTPHandleFunc(handlers.HandleAuthorByID))
	router.HandleFunc("/book", helpers.MakeHTTPHandleFunc(handlers.HandleBook))
	router.HandleFunc("/book/{id}", helpers.MakeHTTPHandleFunc(handlers.HandleBookByID))

	return router
}
