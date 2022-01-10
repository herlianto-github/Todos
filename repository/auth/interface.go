package auth

import "todos/entities"

type AuthInterface interface {
	LoginUser(name, password string) (entities.User, error)
}
