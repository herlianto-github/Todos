package routes

import (
	"todos/delivery/controllers/auth"
	"todos/delivery/controllers/project"
	"todos/delivery/controllers/todo"
	"todos/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(
	e *echo.Echo, authctrl *auth.AuthController, uctrl *user.UsersController, pctrl *project.ProjectsController,
	tdctrl *todo.ToDoController,
) {

	e.Use(middleware.RemoveTrailingSlash())
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
	e.POST("/todos", tdctrl.PostTodoCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/todos/all", tdctrl.GetAllTodoCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/todos/:id", tdctrl.GetTodoCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.PUT("/todos", tdctrl.PutTodoCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.DELETE("/todos", tdctrl.DeleteTodoCtrl(), middleware.JWT([]byte("RAHASIA")))

	// ---------------------------------------------------------------------
	// CRUD Projects
	// ---------------------------------------------------------------------
	e.POST("/projects", pctrl.PostProjectsCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/projects/all", pctrl.GetAllProjectsCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/projects", pctrl.GetProjectsCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.PUT("/projects/:id", pctrl.PutProjectsCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.DELETE("/projects/:id", pctrl.DeleteProjectsCtrl(), middleware.JWT([]byte("RAHASIA")))

}
