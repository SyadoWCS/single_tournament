package server

import (
	"github.com/SyadoWCS/single_tournament/controller"
	"github.com/labstack/echo/v4"
)

func Setup(app *echo.Echo) {
	app.GET("/", controller.Home)
	app.GET("/other", controller.Other)
	app.POST("/api/register", controller.Register)
}
