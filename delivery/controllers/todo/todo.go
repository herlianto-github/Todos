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
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		todos, err := tdcon.Repo.Get(id, userID)

		if todos.ID != 0 && err != nil {
			return c.JSON(
				http.StatusOK, map[string]interface{}{
					"message": "succes",
					"data":    todos,
				},
			)
		}

		return c.JSON(
			http.StatusNotFound, map[string]interface{}{
				"message": common.NewNotFoundResponse(),
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
		if err := c.Bind(&PutTodoReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newTodo := entities.ToDo{
			Task: PutTodoReq.Task,

			Description: PutTodoReq.Description,
		}

		_, err := tdcon.Repo.Update(newTodo, PutTodoReq.ToDoID, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}
