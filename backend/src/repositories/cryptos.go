package repositories

import (
	"backend/src/entities"
	enumhelper "backend/src/enum-helpers"

	"gorm.io/gorm"
)

type CryptoRepository struct {
	db *gorm.DB
}

func NewCryptoRepository(db *gorm.DB) *CryptoRepository {
	return &CryptoRepository{db}
}

func (this CryptoRepository) GetPageOfCryptos(page uint) ([]entities.Crypto, error) {
	var offset uint = 0
	var cryptos []entities.Crypto

	if page > 1 {
		offset = (page - uint(1)) * uint(enumhelper.PageSize)
	}

	result := this.db.Limit(enumhelper.PageSize).Offset(int(offset)).Find(&cryptos)
	if result.Error != nil {
		return cryptos, result.Error
	}

	return cryptos, nil
}

func (this CryptoRepository) GetCryptoById(cryptoId uint) (entities.Crypto, error) {
	var crypto entities.Crypto

	result := this.db.First(&crypto, cryptoId)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return crypto, result.Error
	}

	return crypto, nil
}

func (this CryptoRepository) Vote(cryptoId uint, crypto *entities.Crypto) error {
	result := this.db.Where("id = ?", cryptoId).Updates(crypto)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
