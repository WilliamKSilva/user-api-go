package handlers

import (
	"net/http"

	"github.com/WilliamKSilva/unit-tests-go/application/entities"
	"github.com/WilliamKSilva/unit-tests-go/application/usecases"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase usecases.UserUseCase
}

func Create(c echo.Context, u *UserHandler) error {
	var user entities.User
	err := c.Bind(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u.UserUseCase.Create(&user)
	return c.JSON(http.StatusOK, user)
}
