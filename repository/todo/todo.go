package todo

import (
	"todos/entities"

	"gorm.io/gorm"
)

type ToDoRepository struct {
	db *gorm.DB
}

func NewToDoRepo(db *gorm.DB) *ToDoRepository {
	return &ToDoRepository{db: db}
}

func (tdrep *ToDoRepository) GetAll() ([]entities.ToDo, error) {
	toDos := []entities.ToDo{}
	tdrep.db.Find(&toDos)
	return toDos, nil
}

func (tdrep *ToDoRepository) Get(toDoId int) (entities.ToDo, error) {
	toDo := entities.ToDo{}
	tdrep.db.Find(&toDo, toDoId)
	return toDo, nil
}

func (tdrep *ToDoRepository) Create(toDo entities.ToDo) (entities.ToDo, error) {
	tdrep.db.Save(&toDo)
	return toDo, nil
}

func (tdrep *ToDoRepository) Delete(toDoId int) (entities.ToDo, error) {
	toDo := entities.ToDo{}
	tdrep.db.Find(&toDo, "id=?", toDoId)
	tdrep.db.Delete(&toDo)
	return toDo, nil
}

func (tdrep *ToDoRepository) Update(newToDo entities.ToDo, toDoId int) (entities.ToDo, error) {
	toDo := entities.ToDo{}
	tdrep.db.Find(&toDo, "id=?", toDoId)
	tdrep.db.Model(&toDo).Updates(newToDo)
	return newToDo, nil
}
