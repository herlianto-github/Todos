package todo

import "todos/entities"

type CreateToDoResponseFormat struct {
	Message string          `json:"message"`
	Data    []entities.ToDo `json:"data"`
}

type GetToDoResponseFormat struct {
	Message interface{}     `json:"message"`
	Data    []entities.ToDo `json:"data"`
}
type GetAllToDoResponseFormat struct {
	Message string          `json:"message"`
	Data    []entities.ToDo `json:"data"`
}
type DeleteToDoResponseFormat struct {
	Message string `json:"message"`
}

type PutToDoResponseFormat struct {
	Message string          `json:"message"`
	Data    []entities.ToDo `json:"data"`
}
