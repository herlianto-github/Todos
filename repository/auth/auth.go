package auth

import (
	"todos/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) LoginUser(name, password string) (entities.User, error) {
	var user entities.User
	var err error
	if err = a.db.Where("name = ? AND password = ?", name, password).First(&user).Error; err != nil {
		return user, err
	}

	if err := a.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
