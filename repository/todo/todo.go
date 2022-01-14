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

func (td *ToDoRepository) GetAll(userID int) ([]entities.ToDo, error) {
	toDos := []entities.ToDo{}
	td.db.Where("user_id = ?", userID).Find(&toDos)

	return toDos, nil
}

func (td *ToDoRepository) Get(toDoId, userId int) ([]entities.ToDo, error) {
	toDo := []entities.ToDo{}
	td.db.Where("id = ? AND user_id = ?", toDoId, userId).Find(&toDo)

	return toDo, nil
}

func (td *ToDoRepository) Create(toDo entities.ToDo) (entities.ToDo, error) {
	td.db.Save(&toDo)
	return toDo, nil
}

func (td *ToDoRepository) Delete(toDoId, userID int) (entities.ToDo, error) {
	toDo := entities.ToDo{}
	td.db.Find(&toDo, "id=?", toDoId, userID)
	td.db.Delete(&toDo)
	return toDo, nil
}

func (td *ToDoRepository) Update(newToDo entities.ToDo, toDoId, userId int) ([]entities.ToDo, error) {
	toDo := []entities.ToDo{}

	td.db.Where(
		"id = ? AND user_id = ?", toDoId, userId,
	).Find(&toDo).Save(map[string]interface{}{"task": newToDo.Task, "description": newToDo.Description})

	return toDo, nil
}
