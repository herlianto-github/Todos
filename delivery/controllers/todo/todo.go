package todo

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todos/delivery/common"
	"todos/entities"
	"todos/repository/todo"
)

type ToDoController struct {
	Repo todo.ToDoInterface
}

func NewToDoControllers(torep todo.ToDoInterface) *ToDoController {
	return &ToDoController{Repo: torep}
}

//CreateTodo

func (tdcon ToDoController) PostTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newTodoReq := CreateToDoRequestFormat{}

		if err := c.Bind(&newTodoReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newTodo := entities.ToDo{
			Task:        newTodoReq.Task,
			Description: newTodoReq.Description,
			UserID:      newTodoReq.UserID,
			ProjectID:   newTodoReq.ProjectID,
		}

		_, err := tdcon.Repo.Create(newTodo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}
func (tdcon ToDoController) GetAllTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		userId := GetAllToDoRequestFormat{}

		if err := c.Bind(&userId); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		todos, err := tdcon.Repo.GetAll(userId.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "success",
				"data":    todos,
			},
		)
	}

}
func (tdcon ToDoController) GetTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		ToDoId := GetToDoRequestFormat{}

		if err := c.Bind(&ToDoId); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		todos, err := tdcon.Repo.Get(ToDoId.ToDoID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "success",
				"data":    todos,
			},
		)
	}

}
func (tdcon ToDoController) DeleteTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		ToDoId := DeleteToDoRequestFormat{}

		if err := c.Bind(&ToDoId); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		_, err := tdcon.Repo.Delete(ToDoId.ToDoID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "success",
			},
		)
	}

}
func (tdcon ToDoController) PutTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		PutTodoReq := PutToDoRequestFormat{}

		if err := c.Bind(&PutTodoReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newTodo := entities.ToDo{
			Task: PutTodoReq.Task,

			Description: PutTodoReq.Description,
		}

		_, err := tdcon.Repo.Update(newTodo, PutTodoReq.ToDoID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}
