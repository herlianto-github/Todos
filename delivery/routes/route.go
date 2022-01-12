package routes

import (
	"todos/delivery/controllers/auth"
	"todos/delivery/controllers/todo"
	"todos/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(
	e *echo.Echo, authctrl *auth.AuthController, uctrl *users.UsersController, toctrl *todo.ToDoController,
) {

	// ---------------------------------------------------------------------
	// Login & Register
	// ---------------------------------------------------------------------
	e.POST("/users/register", uctrl.PostUserCtrl())
	e.POST("/users/login", authctrl.LoginAuthCtrl())

	// ---------------------------------------------------------------------
	// CRUD Users
	// ---------------------------------------------------------------------
	e.GET("/users", uctrl.GetUsersCtrl(), middleware.JWT([]byte("RAHASIA")))

	// ---------------------------------------------------------------------
	// CRUD Todo
	// ---------------------------------------------------------------------
	e.POST("/todo", toctrl.PostTodoCtrl())
	e.GET("/todo", toctrl.GetAllTodoCtrl())
	e.GET("/todssso", toctrl.GetTodoCtrl())
	e.PUT("/todo", toctrl.PutTodoCtrl())
	e.DELETE("/todo", toctrl.DeleteTodoCtrl())

}
