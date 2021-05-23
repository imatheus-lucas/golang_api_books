package main

import (
	"github.com/imatheus-lucas/golang_api/database"
	"github.com/imatheus-lucas/golang_api/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
