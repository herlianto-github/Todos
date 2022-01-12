package project

import (
	"fmt"
	"todos/entities"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (prrep *ProjectRepository) GetAll(userId int) ([]entities.Project, error) {
	project := []entities.Project{}
	prrep.db.Where("user_id = ?", userId).Find(&project)
	return project, nil
}

func (prrep *ProjectRepository) Get(projectId int) (entities.Project, error) {
	project := entities.Project{}
	prrep.db.Find(&project, projectId)
	fmt.Println("repo", project)
	return project, nil
}

func (prrep *ProjectRepository) Create(project entities.Project) (entities.Project, error) {
	prrep.db.Save(&project)
	return project, nil
}

func (prrep *ProjectRepository) Delete(projectId int) (entities.Project, error) {
	project := entities.Project{}
	prrep.db.Find(&project, "id=?", projectId)
	prrep.db.Delete(&project)
	return project, nil
}

func (prrep *ProjectRepository) Update(newProject entities.Project, projectId int) (entities.Project, error) {
	project := entities.Project{}
	prrep.db.Find(&project, "id=?", projectId)
	prrep.db.Model(&project).Update("project_name", newProject.ProjectName)
	return newProject, nil
}
