package database

import (
	"time"

	"gorm.io/gorm"
)

type Crypto struct {
	gorm.Model
	ID        int
	Name      string `json:"cryptoName"`
	Code      string `json:"cryptoSimbol"`
	Votes     int    `json:"votes" gorm:"default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	gorm.Model
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
