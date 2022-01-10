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

func (ur *UserRepository) GetAll() ([]entities.User, error) {
	users := []entities.User{}
	ur.db.Find(&users)
	return users, nil
}

func (ur *UserRepository) Get(userId int) (entities.User, error) {
	user := entities.User{}
	ur.db.Find(&user, userId)
	return user, nil
}

func (ur *UserRepository) Create(user entities.User) (entities.User, error) {
	ur.db.Save(&user)
	return user, nil
}

func (ur *UserRepository) Delete(userId int) (entities.User, error) {
	user := entities.User{}
	ur.db.Find(&user, "id=?", userId)
	ur.db.Delete(&user)
	return user, nil
}

func (ur *UserRepository) Update(newUser entities.User, userId int) (entities.User, error) {
	user := entities.User{}
	ur.db.Find(&user, "id=?", userId)
	ur.db.Model(&user).Updates(newUser)
	return newUser, nil
}
