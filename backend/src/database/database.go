package database

import (
	"backend/src/config"
	database "backend/src/database/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func readCryptoFile() ([]database.Crypto, error) {
	rootPath, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/cryptos.json", rootPath)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []database.Crypto{}, errors.New("Error when trying to read the specified file")
	}

	var cryptos []database.Crypto
	err = json.Unmarshal(content, &cryptos)
	if err != nil {
		return []database.Crypto{}, errors.New("Error when trying to parse the specified file")
	}

	return cryptos, nil
}

func seedCryptoTable(db *gorm.DB) {
	data, err := readCryptoFile()
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&data)
}

func RunMigrationsAndSeeds(db *gorm.DB) {
	if err := db.AutoMigrate(&database.Crypto{}); err == nil && db.Migrator().HasTable(&database.Crypto{}) {
		if err := db.First(&database.Crypto{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seedCryptoTable(db)
		}
	}
}

func Connect() (*gorm.DB, error) {
	URL := config.DB_URL
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
