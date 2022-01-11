package project

import "todos/entities"

type ProjectInterface interface {
	GetAll() ([]entities.Project, error)
	Get(projectID int) (entities.Project, error)
	Create(project entities.Project) (entities.Project, error)
	Delete(projectID int) (entities.Project, error)
	Update(newProject entities.Project, bookId int) (entities.Project, error)
}
