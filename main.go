package main

import (
	"github.com/WilliamKSilva/user-api-go/application/repositories"
	usecases "github.com/WilliamKSilva/user-api-go/application/usecases/user"
	"github.com/WilliamKSilva/user-api-go/handlers"
	"github.com/WilliamKSilva/user-api-go/infra/postgres"
	"github.com/labstack/echo/v4"
)

var db = postgres.Init()

func main() {
	e := echo.New()

	userHandler := setupUserHandler()
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.FindAllUser)
	e.GET("/users/:id", userHandler.FindUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	println("Server is Running!")
	e.Logger.Fatal(e.Start(":3000"))
}

func setupUserHandler() *handlers.UserHandler {
	userHandler := handlers.NewUserHandler()
	userRepository := repositories.UserRepositoryDb{Db: db}
	userHandler.UserUseCase = usecases.UserUseCase{UserRepository: userRepository}
	return userHandler
}
