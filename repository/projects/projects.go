package projects

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

func (pr *ProjectRepository) GetAll() ([]entities.Projects, error) {
	projects := []entities.Projects{}
	pr.db.Find(&projects)
	return projects, nil
}

func (pr *ProjectRepository) Get(projectId int) (entities.Projects, error) {
	project := entities.Projects{}
	pr.db.Find(&project, projectId)
	return project, nil
}

func (pr *ProjectRepository) Create(project entities.Projects) (entities.Projects, error) {
	pr.db.Save(&project)
	return project, nil
}

func (pr *ProjectRepository) Delete(projectId int) (entities.Projects, error) {
	project := entities.Projects{}
	pr.db.Find(&project, "id=?", projectId)
	pr.db.Delete(&project)
	return project, nil
}

func (pr *ProjectRepository) Update(newProject entities.Projects, projectId int) (entities.Projects, error) {
	project := entities.Projects{}
	pr.db.Find(&project, "id=?", projectId)
	pr.db.Model(&project).Updates(newProject)
	return newProject, nil
}
