package main

import (
	"backend/src/config"
	"backend/src/database"
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.InitialConfigs()

	router := router.GetRouter()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	database.RunMigrationsAndSeeds(db)

	fmt.Printf("Server listening on port: %d!", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), router))
}
