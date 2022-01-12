package todo

import (
	"fmt"
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

func (td *ToDoRepository) Get(toDoId int) (entities.ToDo, error) {
	toDo := entities.ToDo{}
	td.db.Find(&toDo, toDoId)
	return toDo, nil
}

func (td *ToDoRepository) Create(toDo entities.ToDo) (entities.ToDo, error) {
	td.db.Save(&toDo)
	return toDo, nil
}

func (td *ToDoRepository) Delete(toDoId int) (entities.ToDo, error) {
	toDo := entities.ToDo{}
	td.db.Find(&toDo, "id=?", toDoId)
	td.db.Delete(&toDo)
	return toDo, nil
}

func (td *ToDoRepository) Update(newToDo entities.ToDo, toDoId int) (entities.ToDo, error) {
	toDo := entities.To_Do{}
	fmt.Println(newToDo.Task)
	td.db.Find(&toDo, "id=?", toDoId)
	td.db.Model(&toDo).Updates(
		map[string]interface{}{
			"task": newToDo.Task, "description": newToDo.Description,
		},
	)

	return newToDo, nil
}
