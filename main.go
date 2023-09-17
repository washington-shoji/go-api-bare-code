package main

import (
	"log"

	"github.com/washington-shoji/gobare/databases"
	"github.com/washington-shoji/gobare/server"
)

func main() {

	store, err := databases.PostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := server.Server(":3030")
	server.Run()
}
