package main

import (
	"backend/src/config"
	"backend/src/database"
	"backend/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type teste struct {
	vote int
}

func main() {
	config.InitialConfigs()

	router := router.GetRouter()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	database.RunMigrationsAndSeeds(db)

	fmt.Printf("Server listening on port: %d!\n", config.PORT)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	httpHandler := cors.Default().Handler(router)
	xablau := c.Handler(httpHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), xablau))
}
