package todo

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todos/delivery/common"
	"todos/entities"
	"todos/repository/to_do"
)

type ToDoController struct {
	Repo to_do.To_Do
}

func NewToDoControllers(torep to_do.To_Do) *ToDoController {
	return &ToDoController{Repo: torep}
}

//CreateTodo

func (tocon ToDoController) PostTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newTodoReq := CreateToDoRequestFormat{}

		if err := c.Bind(&newTodoReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newTodo := entities.To_Do{
			Task:        newTodoReq.Task,
			Description: newTodoReq.Description,
			UserID:      newTodoReq.UserID,
			ProjectID:   newTodoReq.ProjectID,
		}

		_, err := tocon.Repo.Create(newTodo)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}
func (tocon ToDoController) GetAllTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		userId := GetAllToDoRequestFormat{}

		if err := c.Bind(&userId); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		todos, err := tocon.Repo.GetAll(userId.UserID)
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
func (tocon ToDoController) GetTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		ToDoId := GetToDoRequestFormat{}

		if err := c.Bind(&ToDoId); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		todos, err := tocon.Repo.Get(ToDoId.ToDoID)
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
func (tocon ToDoController) DeleteTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		ToDoId := DeleteToDoRequestFormat{}

		if err := c.Bind(&ToDoId); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		_, err := tocon.Repo.Delete(ToDoId.ToDoID)
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
func (tocon ToDoController) PutTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		PutTodoReq := PutToDoRequestFormat{}

		if err := c.Bind(&PutTodoReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newTodo := entities.To_Do{
			Task:        PutTodoReq.Task,
			Status:      PutTodoReq.Description,
			Description: PutTodoReq.Description,
		}

		_, err := tocon.Repo.Update(newTodo, PutTodoReq.ToDoID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}
