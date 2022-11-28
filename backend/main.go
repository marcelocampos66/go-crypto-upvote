package main

import (
	"backend/src/config"
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.InitialConfigs()

	router := router.GetRouter()

	fmt.Printf("Server listening on port: %d!", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), router))
}
