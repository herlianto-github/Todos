package projects

import (
	"todos/entities"

	"github.com/labstack/gommon/log"
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
	if err := pr.db.Find(&projects).Error; err != nil {
		log.Warn("Found database error", err)
		return nil, err
	}
	return projects, nil
}

func (pr *ProjectRepository) Get(projectId int) (entities.Projects, error) {
	project := entities.Projects{}
	if err := pr.db.Find(&project, projectId).Error; err != nil {
		log.Warn("Found database error", err)
		return project, err
	}
	return project, nil
}

func (pr *ProjectRepository) Create(project entities.Projects) (entities.Projects, error) {
	if err := pr.db.Save(&project).Error; err != nil {
		return project, err
	}
	return project, nil
}

func (pr *ProjectRepository) Delete(projectId int) (entities.Projects, error) {
	project := entities.Projects{}
	if err := pr.db.Find(&project, "id=?", projectId).Error; err != nil {
		return project, err
	}
	if err := pr.db.Delete(&project).Error; err != nil {
		return project, err
	}
	return project, nil
}

func (pr *ProjectRepository) Update(newProject entities.Projects, projectId int) (entities.Projects, error) {
	project := entities.Projects{}
	if err := pr.db.Find(&project, "id=?", projectId).Error; err != nil {
		return project, err
	}
	if err := pr.db.Model(&project).Updates(newProject).Error; err != nil {
		return project, err
	}
	return newProject, nil
}
