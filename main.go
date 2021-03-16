package main

import (
	"log"

	"github.com/FrancoCastro-94/Go-Mongo-practice/db"
	"github.com/FrancoCastro-94/Go-Mongo-practice/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No connection to BD")
		return
	}
	handlers.Handlers()
}
