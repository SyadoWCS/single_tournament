package server

import (
	"github.com/SyadoWCS/single_tournament/controller"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.GET("/", controller.Home)
	e.GET("/other", controller.Other)
	e.POST("/api/register", controller.Register)
	e.POST("/api/login", controller.Login)
	e.GET("/api/user", controller.User)
	e.GET("/api/logout", controller.Logout)
}
