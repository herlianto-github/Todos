package user

import "todos/entities"

type RegisterUserResponseFormat struct {
	Message string          `json:"message"`
	Data    []entities.User `json:"data"`
}

type LoginUserResponseFormat struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type GetUsersResponseFormat struct {
	Message string          `json:"message"`
	Data    []entities.User `json:"data"`
}
