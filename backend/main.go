package main

import (
	"backend/src/config"
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

	fmt.Printf("Server listening on port: %d!\n", config.PORT)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		Debug:            true,
	})
	corsHandler := cors.Default().Handler(router)
	httpHandler := c.Handler(corsHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), httpHandler))
}
