package server

import (
	"github.com/SyadoWCS/single_tournament/controller"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.Pre(controller.MethodOverride)
	e.GET("/", controller.Home)
	e.POST("/api/register", controller.Register)
	e.POST("/api/login", controller.Login)
	e.GET("/api/user", controller.User)
	e.GET("/api/logout", controller.Logout)
	e.GET("/tournament/list", controller.TournamentList)
	e.GET("/tournament/new", controller.TournamentNew)
	e.POST("/tournament/create", controller.TournamentCreate)
	e.GET("/tournament/edit/:id", controller.TournamentEdit)
	e.PUT("/tournament/update/:id", controller.TournamentUpdate)
	e.DELETE("/tournament/delete/:id", controller.TournamentDelete)
	e.GET("/entry/list/:tournament_id", controller.EntryList)
	e.GET("/entry/new/:tournament_id", controller.EntryNew)
	e.POST("/entry/create/:tournament_id", controller.EntryCreate)
	e.DELETE("/entry/delete/:tournament_id/:entry_id", controller.EntryDelete)
}
