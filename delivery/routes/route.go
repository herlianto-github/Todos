package routes

import (
	"todos/delivery/controllers/auth"
	"todos/delivery/controllers/project"
	"todos/delivery/controllers/todo"
	"todos/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, authctrl *auth.AuthController, uctrl *user.UsersController, pctrl *project.ProjectsController, tdctrl *todo.ToDoController) {

	// ---------------------------------------------------------------------
	// Login & Register
	// ---------------------------------------------------------------------
	e.POST("/users/register", uctrl.PostUserCtrl())
	e.POST("/users/login", authctrl.LoginAuthCtrl())

	// ---------------------------------------------------------------------
	// CRUD Users
	// ---------------------------------------------------------------------
	e.GET("/users", uctrl.GetUsersCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/users/:id", uctrl.GetUserCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.PUT("/users/:id", uctrl.EditUserCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.DELETE("/users/:id", uctrl.DeleteUserCtrl(), middleware.JWT([]byte("RAHASIA")))

	// ---------------------------------------------------------------------
	// CRUD Projects
	// ---------------------------------------------------------------------

	// ---------------------------------------------------------------------
	// CRUD Todos
	// ---------------------------------------------------------------------
	e.POST("/todo", tdctrl.PostTodoCtrl())
	e.GET("/todo/all", tdctrl.GetAllTodoCtrl())
	e.GET("/todo", tdctrl.GetTodoCtrl())
	e.PUT("/todo", tdctrl.PutTodoCtrl())
	e.DELETE("/todo", tdctrl.DeleteTodoCtrl())

	// ---------------------------------------------------------------------
	// CRUD Projects
	// ---------------------------------------------------------------------
	e.POST("/project", pctrl.PostProjectsCtrl())
	e.GET("/project/all", pctrl.GetAllProjectsCtrl())
	e.GET("/project", pctrl.GetProjectsCtrl())
	e.PUT("/project", pctrl.PutProjectsCtrl())
	e.DELETE("/project", pctrl.DeleteProjectsCtrl())

}
