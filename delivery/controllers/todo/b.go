package todo

import "todos/entities"

type GetToDosResponseFormat struct {
	Message string          `json:"message"`
	Data    []entities.ToDo `json:"data"`
}
