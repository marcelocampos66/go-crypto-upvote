package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DB_URL     = ""
	PORT       = 0
	JWT_SECRET = ""
	MARKET_API = ""
)

func loadEnviromentVars() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		PORT = 3001
	}

	DB_URL = os.Getenv("DB_URL")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	MARKET_API = os.Getenv("MARKET_API")
}

func InitialConfigs() {
	fmt.Println("Running initial configs")

	loadEnviromentVars()
}
