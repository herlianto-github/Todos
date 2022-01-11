package project

import (
	"todos/entities"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (prr *ProjectRepository) GetAll() ([]entities.Project, error) {
	project := []entities.Project{}
	prr.db.Find(&project)
	return project, nil
}

func (prr *ProjectRepository) Get(projectId int) (entities.Project, error) {
	project := entities.Project{}
	prr.db.Find(&project, projectId)
	return project, nil
}

func (prr *ProjectRepository) Create(project entities.Project) (entities.Project, error) {
	prr.db.Save(&project)
	return project, nil
}

func (prr *ProjectRepository) Delete(projectId int) (entities.Project, error) {
	project := entities.Project{}
	prr.db.Find(&project, "id=?", projectId)
	prr.db.Delete(&project)
	return project, nil
}

func (prr *ProjectRepository) Update(newProject entities.Project, projectId int) (entities.Project, error) {
	project := entities.Project{}
	prr.db.Find(&project, "id=?", projectId)
	prr.db.Model(&project).Updates(newProject)
	return newProject, nil
}
