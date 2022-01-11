package to_do

import "todos/entities"

type To_DoInterface interface {
	GetAll() ([]entities.To_Do, error)
	Get(toDoId int) (entities.To_Do, error)
	Create(todo entities.To_Do) (entities.To_Do, error)
	Delete(toDoId int) (entities.To_Do, error)
	Update(newToDo entities.To_Do, toDoId int) (entities.To_Do, error)
}
