package repositories

import (
	"backend/src/config"
	"backend/src/database"
	"backend/src/entities"
	enumhelper "backend/src/enum-helpers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type ICryptoRepository interface {
	GetPageOfCryptos(page uint) ([]entities.Crypto, error)
	GetCryptoById(cryptoId uint) (entities.Crypto, error)
	Vote(cryptoId uint, crypto *entities.Crypto) error
	GetCryptoLastQuotation(crypto entities.Crypto, wg *sync.WaitGroup) entities.Crypto
}

type CryptoRepository struct{}

func (this CryptoRepository) GetPageOfCryptos(page uint) ([]entities.Crypto, error) {
	var offset uint = 0
	var cryptos []entities.Crypto

	db, err := database.Connect()
	if err != nil {
		return cryptos, err
	}

	if page > 1 {
		offset = (page - uint(1)) * uint(enumhelper.PageSize)
	}

	result := db.Limit(enumhelper.PageSize).Offset(int(offset)).Order("id asc").Find(&cryptos)
	if result.Error != nil {
		return cryptos, result.Error
	}

	return cryptos, nil
}

func (this CryptoRepository) GetCryptoById(cryptoId uint) (entities.Crypto, error) {
	var crypto entities.Crypto

	db, err := database.Connect()
	if err != nil {
		return crypto, err
	}

	result := db.First(&crypto, cryptoId)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return crypto, result.Error
	}

	return crypto, nil
}

func (this CryptoRepository) Vote(cryptoId uint, crypto *entities.Crypto) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	result := db.Where("id = ?", cryptoId).Updates(crypto)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (this CryptoRepository) GetCryptoLastQuotation(crypto entities.Crypto, wg *sync.WaitGroup) entities.Crypto {
	defer wg.Done()
	var HOST string = config.MARKET_API
	var url string = fmt.Sprintf("%s/api/%s/ticker", HOST, crypto.Code)

	response, err := http.Get(url)
	if err != nil {
		return crypto
	}
	defer response.Body.Close()

	type quotation struct {
		High string `json:"high,omitempty"`
		Low  string `json:"low,omitempty"`
		Vol  string `json:"vol,omitempty"`
		Last string `json:"last,omitempty"`
		Buy  string `json:"buy,omitempty"`
		Sell string `json:"sell,omitempty"`
		Open string `json:"open,omitempty"`
	}
	type apiResponse struct {
		Ticker quotation `json:"ticker,omitempty"`
	}
	body := apiResponse{}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return crypto
	}

	err = json.Unmarshal(responseBody, &body)
	if err != nil {
		return crypto
	}

	crypto.Quotation = body.Ticker.Last

	return crypto
}
