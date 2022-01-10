package auth

import "todos/entities"

type Auth interface {
	LoginUser(name, password string) (entities.User, error)
}
