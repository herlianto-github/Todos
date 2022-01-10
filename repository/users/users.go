package users

import (
	"todos/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// /users GET
func (usrep *UserRepository) Gets() ([]entities.User, error) {

	users := []entities.User{}
	usrep.db.Find(&users)

	return users, nil
}
