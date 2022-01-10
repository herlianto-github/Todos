package users

import "todos/entities"

type UsersInterface interface {
	Gets() ([]entities.User, error)
}
