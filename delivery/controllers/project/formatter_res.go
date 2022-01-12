package project

import "todos/entities"

type CreateProjectResponseFormat struct {
	Message string             `json:"message"`
	Data    []entities.Project `json:"data"`
}

type GetProjectResponseFormat struct {
	Message string             `json:"message"`
	Data    []entities.Project `json:"data"`
}
type GetAllProjectResponseFormat struct {
	Message string             `json:"message"`
	Data    []entities.Project `json:"data"`
}
type DeleteProjectResponseFormat struct {
	Message string `json:"message"`
}

type PutProjectResponseFormat struct {
	Message string             `json:"message"`
	Data    []entities.Project `json:"data"`
}
