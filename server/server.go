package server

import (
	"log"
	"net/http"

	"github.com/washington-shoji/gobare/routers"
)

type APIServer struct {
	ListenAddress string
}

func Server(listenAddress string) *APIServer {
	return &APIServer{
		ListenAddress: listenAddress,
	}
}

func (server *APIServer) Run() {

	log.Println("JSON API server on port: ", server.ListenAddress)
	router := routers.SetupRoutes()

	http.ListenAndServe(server.ListenAddress, router)
}
