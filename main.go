package main

import (
	"log"

	"github.com/washington-shoji/gobare/databases"
	"github.com/washington-shoji/gobare/server"
)

func main() {

	_, err := databases.PostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := databases.Init(); err != nil {
		log.Fatal(err)
	}

	server := server.Server(":3030")
	server.Run()
}
