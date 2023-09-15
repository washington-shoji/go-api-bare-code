package main

import "github.com/washington-shoji/gobare/server"

func main() {

	server := server.Server(":3030")
	server.Run()
}
