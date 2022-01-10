package projects

import "todos/entities"

type Projects interface {
	GetAll() ([]entities.Projects, error)
	Get(projectID int) (entities.Projects, error)
	Create(project entities.Projects) (entities.Projects, error)
	Delete(projectID int) (entities.Projects, error)
	Update(newProject entities.Projects, bookId int) (entities.Projects, error)
}
