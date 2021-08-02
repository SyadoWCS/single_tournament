package main

import (
	"net/http"

	"github.com/SyadoWCS/single_tournament/database"
	"github.com/SyadoWCS/single_tournament/hello"
	"github.com/SyadoWCS/single_tournament/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// GORMでDB接続
	database.Connetct()

	// Echoのインスタンス作る
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	server.Setup(e)

	// サーバー起動
	e.Start(":80")
}
