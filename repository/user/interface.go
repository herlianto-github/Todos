package user

import "todos/entities"

type UserInterface interface {
	GetAll() ([]entities.User, error)
	Get(userId int) (entities.User, error)
	Create(newUser entities.User) (entities.User, error)
	Update(updateUser entities.User, userId int) (entities.User, error)
	Delete(userId int) (entities.User, error)
}
