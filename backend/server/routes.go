package server

import (
	"github.com/SyadoWCS/single_tournament/controller"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	e.GET("/", controller.Home)
	e.POST("/api/register", controller.Register)
	e.POST("/api/login", controller.Login)
	e.GET("/api/user", controller.User)
	e.GET("/api/logout", controller.Logout)
	e.GET("/api/tournament/list", controller.TournamentList)
	e.GET("/api/tournament/new", controller.TournamentNew)
	e.POST("/api/tournament/create", controller.TournamentCreate)
	e.GET("/api/tournament/edit/:id", controller.TournamentEdit)
	e.POST("/api/tournament/update/:id", controller.TournamentUpdate)
	e.GET("/api/tournament/delete/:id", controller.TournamentDelete)
	e.GET("/api/entry/list/:id", controller.EntryList)
	e.GET("/api/entry/new/:tournament_id", controller.EntryNew)
	e.POST("/api/entry/create/:tournament_id", controller.EntryCreate)
	e.GET("/api/entry/delete/:tournament_id/:entry_id", controller.EntryDelete)
}
