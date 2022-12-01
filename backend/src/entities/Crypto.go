package entities

import (
	"time"
)

type Crypto struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"cryptoName"`
	Code      string    `json:"cryptoSimbol"`
	Votes     int       `json:"votes"`
	ImgUrl    string    `json:"imageUrl"`
	Quotation string    `json:"quotation"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
