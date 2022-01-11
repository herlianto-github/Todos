package todo

import (
	"net/http"
	"todos/delivery/common"
	"todos/entities"
	"todos/repository/todo"

	"github.com/labstack/echo/v4"
)

type TodosController struct {
	Repo todo.ToDoInterface
}

func NewTodosControllers(tdrep todo.ToDoInterface) *TodosController {
	return &TodosController{Repo: tdrep}
}

// POST /todos/register
func (tdcon TodosController) PostToDoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newToDoReq := CreateToDoRequestFormat{}

		if err := c.Bind(&newToDoReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newToDo := entities.ToDo{
			ProjectID:   newToDoReq.ProjectID,
			UserID:      newToDoReq.UserID,
			Task:        newToDoReq.Task,
			Status:      newToDoReq.Status,
			Description: newToDoReq.Description,
		}

		_, err := tdcon.Repo.Create(newToDo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}

// GET /todos
func (tdcon TodosController) GetTodosCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		todos, err := tdcon.Repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		response := GetToDosResponseFormat{
			Message: "Successful Operation",
			Data:    todos,
		}

		return c.JSON(http.StatusOK, response)
	}
}
