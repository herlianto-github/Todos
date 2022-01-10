package main

import (
	"fmt"
	"log"
	"todos/configs"
	"todos/delivery/controllers/users"
	"todos/delivery/routes"
	userRepo "todos/repository/users"
	"todos/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello Todos")
	config := configs.GetConfig()
	db := utils.InitDB(config)

	e := echo.New()

	userRepo := userRepo.NewUsersRepo(db)
	userCtrl := users.NewUsersControllers(userRepo)

	routes.RegisterPath(e, userCtrl)

	address := fmt.Sprintf("localhost:%d", config.Port)
	log.Fatal(e.Start(address))
}
