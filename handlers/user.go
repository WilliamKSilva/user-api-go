package handlers

import (
	"fmt"
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

func (u *UserHandler) FindAllUser(c echo.Context) error {
	users := u.UserUseCase.FindAll()

	return c.JSON(http.StatusOK, users)
}

func (u *UserHandler) FindUser(c echo.Context) error {
	id := c.Param("id")

	converted_id, convert_error := strconv.ParseUint(id, 0, 64)

	fmt.Println(converted_id)

	if convert_error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, convert_error.Error())
	}

	user, err := u.UserUseCase.Find(uint(converted_id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	converted_id, convert_error := strconv.ParseUint(id, 0, 64)
	if convert_error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, convert_error.Error())
	}

	err := u.UserUseCase.Delete(uint(converted_id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	var user entities.User
	id := c.Param("id")
	err := c.Bind(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	converted_id, convert_error := strconv.ParseUint(id, 0, 64)
	if convert_error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, convert_error.Error())
	}

	err_update := u.UserUseCase.Update(uint(converted_id), &user)

	if err_update != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_update.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
