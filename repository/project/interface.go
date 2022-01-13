package project

import "todos/entities"

type ProjectInterface interface {
	GetAll(userId int) ([]entities.Project, error)
	Get(projectID, userId int) (entities.Project, error)
	Create(project entities.Project) (entities.Project, error)
	Update(newProject entities.Project, projectId, Userid int) (entities.Project, error)
	Delete(projectID, userId int) (entities.Project, error)
}
