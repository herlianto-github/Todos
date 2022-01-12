package project

import "todos/entities"

type GetProjectsResponseFormat struct {
	Message string             `json:"message"`
	Data    []entities.Project `json:"data"`
}
