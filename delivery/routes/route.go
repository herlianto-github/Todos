package routes

import (
	"todos/delivery/controllers/auth"
	"todos/delivery/controllers/project"
	"todos/delivery/controllers/todo"
	"todos/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, authctrl *auth.AuthController, uctrl *user.UsersController, pctrl *project.ProjectsController, tdctrl *todo.TodosController) {

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
	// CRUD Projects
	// ---------------------------------------------------------------------
	e.POST("/projects/register", pctrl.PostToDoCtrl())
	e.GET("/projects", pctrl.GetProjectsCtrl(), middleware.JWT([]byte("RAHASIA")))

	// ---------------------------------------------------------------------
	// CRUD Todos
	// ---------------------------------------------------------------------
	e.POST("/todos/register", tdctrl.PostToDoCtrl())
	e.GET("/todos", tdctrl.GetTodosCtrl(), middleware.JWT([]byte("RAHASIA")))

}
