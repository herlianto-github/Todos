package users

import "todos/entities"

type UsersInterface interface {
	GetAll() ([]entities.User, error)
	Get(userId int) (entities.User, error)
	Create(user entities.User) (entities.User, error)
	Delete(userId int) (entities.User, error)
	Update(nweUser entities.User, userId int) (entities.User, error)
}
