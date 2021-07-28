package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Other(c echo.Context) error {
	return c.String(http.StatusOK, "Other Controller")
}
