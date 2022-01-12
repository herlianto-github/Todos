package main

import (
	"fmt"
	"log"
	"todos/configs"
	"todos/delivery/controllers/auth"
	"todos/delivery/controllers/todo"
	"todos/delivery/controllers/users"
	"todos/delivery/routes"
	authRepo "todos/repository/auth"
	todoRepo "todos/repository/to_do"
	userRepo "todos/repository/users"
	"todos/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello Todos")
	config := configs.GetConfig("")
	db := utils.InitDB(config)

	e := echo.New()

	authRepo := authRepo.NewAuthRepo(db)
	authCtrl := auth.NewAuthControllers(authRepo)

	userRepo := userRepo.NewUsersRepo(db)
	userCtrl := users.NewUsersControllers(userRepo)

	todoRepo := todoRepo.NewTo_DoRepo(db)
	todoCtrl := todo.NewToDoControllers(todoRepo)

	routes.RegisterPath(e, authCtrl, userCtrl, todoCtrl)

	address := fmt.Sprintf("localhost:%d", config.Port)
	log.Fatal(e.Start(address))
}
