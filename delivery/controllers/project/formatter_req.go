package project

import "todos/entities"

type CreateProjectRequestFormat struct {
	ProjectName string `json:"projectname" form:"projectname"`
	Todo        []entities.ToDo
}

type PutProjectRequestFormat struct {
	ProjectName string `json:"projectname" form:"projectname"`
	Todo        []entities.ToDo
}
