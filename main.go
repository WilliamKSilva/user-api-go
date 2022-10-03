package main

import (
	"net/http"

	"github.com/WilliamKSilva/unit-tests-go/application/repositories"
	"github.com/WilliamKSilva/unit-tests-go/application/usecases"
	"github.com/WilliamKSilva/unit-tests-go/handlers"
	"github.com/WilliamKSilva/unit-tests-go/infra/postgres"
	"github.com/labstack/echo/v4"
)

var db = postgres.Init()

func createUser(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func main() {
	e := echo.New()

	userHandler := setupUserHandler()

	e.POST("/users", userHandler.CreateUser)

	println("Server is Running!")
	e.Logger.Fatal(e.Start(":3000"))
}

func setupUserHandler() *handlers.UserHandler {
	userHandler := handlers.NewUserHandler()
	userRepository := repositories.UserRepositoryDb{Db: db}
	userHandler.UserUseCase = usecases.UserUseCase{UserRepository: userRepository}
	return userHandler
}
