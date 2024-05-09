package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	dbpostgres "github.com/kolaboradev/inventory/src/database/postgres"
	httpServer "github.com/kolaboradev/inventory/src/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())
	}
	db, err := dbpostgres.NewDB()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	fmt.Println("Sukses connect")

	serverHttp := httpServer.NewServer(db)

	serverHttp.Listen()
}
