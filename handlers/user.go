package handlers

import (
	"net/http"
	"strconv"

	"github.com/WilliamKSilva/unit-tests-go/application/entities"
	"github.com/WilliamKSilva/unit-tests-go/application/usecases"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase usecases.UserUseCase
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var user entities.User
	err := c.Bind(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u.UserUseCase.Create(&user)
	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) FindUser(c echo.Context) (*entities.User, error) {
	id := c.Param("id")

	converted_id, convert_error := strconv.ParseUint(id, 0, 64)
	if convert_error != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, convert_error.Error())
	}

	user, err := u.UserUseCase.Find(uint(converted_id))

	if err != nil {
		return user, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return user, nil
}
