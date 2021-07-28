package main

import (
	"github.com/SyadoWCS/single_tournament/database"
	"github.com/SyadoWCS/single_tournament/server"
	"github.com/labstack/echo/v4"
)

func main() {
	// GORMでDB接続
	database.Connetct()

	// Echoのインスタンス作る
	app := echo.New()
	server.Setup(app)

	// サーバー起動
	app.Start(":80")
}
