package controller

import (
	"net/http"

	"github.com/SyadoWCS/single_tournament/model"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func Register(c echo.Context) error {
	var user model.User

	user.FirstName = "Self"
	user.LastName = "None"
	user.Email = "selfnote@example.com"
	user.Password = "pass"
	return c.JSON(http.StatusOK, user)
}
