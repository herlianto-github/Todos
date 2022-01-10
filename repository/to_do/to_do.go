package to_do

import (
	"todos/entities"

	"gorm.io/gorm"
)

type To_DoRepository struct {
	db *gorm.DB
}

func NewTo_DoRepo(db *gorm.DB) *To_DoRepository {
	return &To_DoRepository{db: db}
}

func (td *To_DoRepository) GetAll() ([]entities.To_Do, error) {
	to_dos := []entities.To_Do{}
	td.db.Find(&to_dos)
	return to_dos, nil
}
