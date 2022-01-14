package todo

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := uint(claims["userid"].(float64))

		newTodo := entities.ToDo{
			Task:        newTodoReq.Task,
			UserID:      userID,
			Description: newTodoReq.Description,
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

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		todos, err := tdcon.Repo.GetAll(userID)

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
		id, _ := strconv.Atoi(c.Param("id"))

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		todos, _ := tdcon.Repo.Get(id, userID)

		if len(todos) == 0 {
			return c.JSON(
				http.StatusNotFound, map[string]interface{}{
					"message": "Task not found",
				},
			)
		}

		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "succes",
				"data":    todos,
			},
		)

	}

}
func (tdcon ToDoController) DeleteTodoCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		var err error
		id, err := strconv.Atoi(c.Param("id"))
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		_, err = tdcon.Repo.Delete(id, userID)
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
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		err := c.Bind(&PutTodoReq)

		newTodo := entities.ToDo{
			ID:          PutTodoReq.ToDoID,
			Task:        PutTodoReq.Task,
			Description: PutTodoReq.Description,
		}
		if PutTodoReq.ToDoID < 1 || err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		_, err = tdcon.Repo.Update(newTodo, PutTodoReq.ToDoID, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}
