package repositories

import (
	"backend/src/database"
	"backend/src/entities"
	enumhelper "backend/src/enum-helpers"
)

type ICryptoRepository interface {
	GetPageOfCryptos(page uint) ([]entities.Crypto, error)
	GetCryptoById(cryptoId uint) (entities.Crypto, error)
	Vote(cryptoId uint, crypto *entities.Crypto) error
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

	result := db.Limit(enumhelper.PageSize).Offset(int(offset)).Find(&cryptos)
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
