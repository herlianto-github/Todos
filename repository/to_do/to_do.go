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

func (tdr *To_DoRepository) GetAll() ([]entities.To_Do, error) {
	to_Dos := []entities.To_Do{}
	tdr.db.Find(&to_Dos)
	return to_Dos, nil
}

func (tdr *To_DoRepository) Get(to_DoId int) (entities.To_Do, error) {
	to_Do := entities.To_Do{}
	tdr.db.Find(&to_Do, to_DoId)
	return to_Do, nil
}

func (tdr *To_DoRepository) Create(to_Do entities.To_Do) (entities.To_Do, error) {
	tdr.db.Save(&to_Do)
	return to_Do, nil
}

func (tdr *To_DoRepository) Delete(to_DoId int) (entities.To_Do, error) {
	to_Do := entities.To_Do{}
	tdr.db.Find(&to_Do, "id=?", to_DoId)
	tdr.db.Delete(&to_Do)
	return to_Do, nil
}

func (tdr *To_DoRepository) Update(newTo_Do entities.To_Do, to_DoId int) (entities.To_Do, error) {
	to_Do := entities.To_Do{}
	tdr.db.Find(&to_Do, "id=?", to_DoId)
	tdr.db.Model(&to_Do).Updates(newTo_Do)
	return newTo_Do, nil
}
