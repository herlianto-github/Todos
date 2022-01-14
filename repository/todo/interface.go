package todo

import "todos/entities"

type ToDoInterface interface {
	GetAll(userId int) ([]entities.ToDo, error)
	Get(toDoId, userId int) ([]entities.ToDo, error)
	Create(todo entities.ToDo) (entities.ToDo, error)
	Delete(toDoId, userId int) (entities.ToDo, error)
	Update(newToDo entities.ToDo, toDoId, userID int) ([]entities.ToDo, error)
}
