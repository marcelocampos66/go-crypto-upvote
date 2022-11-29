package repositories

import (
	"backend/src/database"
	"backend/src/entities"
)

type IUserRepository interface {
	Create(user *entities.User) (int, error)
	GetUserByEmail(email string) (entities.User, error)
}

type UserRepository struct{}

func (this UserRepository) Create(user *entities.User) (int, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}

	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.Id, nil
}

func (this UserRepository) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User

	db, err := database.Connect()
	if err != nil {
		return user, err
	}

	result := db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
