package todo

import "todos/entities"

type ToDoInterface interface {
	GetAll() ([]entities.ToDo, error)
	Get(toDoId int) (entities.ToDo, error)
	Create(todo entities.ToDo) (entities.ToDo, error)
	Delete(toDoId int) (entities.ToDo, error)
	Update(newToDo entities.ToDo, toDoId int) (entities.ToDo, error)
}
