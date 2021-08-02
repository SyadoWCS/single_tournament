package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func TournamentList(c echo.Context) error {
	return c.String(http.StatusOK, "Tournament List")
}

func TournamentCreate(c echo.Context) error {
	return c.String(http.StatusOK, "Tournament Create")
}

func TournamentDelete(c echo.Context) error {
	return c.String(http.StatusOK, "Tournament Delete")
}
