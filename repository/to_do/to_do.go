package to_do

import (
	"gorm.io/gorm"
)

type To_DoRepository struct {
	db *gorm.DB
}

func NewTo_DoRepo(db *gorm.DB) *To_DoRepository {
	return &To_DoRepository{db: db}
}
