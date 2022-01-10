package projects

import (
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}
