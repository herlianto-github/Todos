package main

import (
	"fmt"
	"log"
	"todos/configs"
	"todos/delivery/controllers/auth"
	"todos/delivery/controllers/project"
	"todos/delivery/controllers/todo"
	"todos/delivery/controllers/user"
	"todos/delivery/routes"
	authRepo "todos/repository/auth"
	projRepo "todos/repository/project"
	todoRepo "todos/repository/todo"
	userRepo "todos/repository/user"
	"todos/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	e := echo.New()

	authRepo := authRepo.NewAuthRepo(db)
	authCtrl := auth.NewAuthControllers(authRepo)

	userRepo := userRepo.NewUsersRepo(db)
	userCtrl := user.NewUsersControllers(userRepo)

	projRepo := projRepo.NewProjectRepo(db)
	projCtrl := project.NewProjectsControllers(projRepo)

	todoRepo := todoRepo.NewToDoRepo(db)
	todoCtrl := todo.NewToDoControllers(todoRepo)

	routes.RegisterPath(e, authCtrl, userCtrl, projCtrl, todoCtrl)

	address := fmt.Sprintf(":%d", config.Port)
	log.Fatal(e.Start(address))
}
