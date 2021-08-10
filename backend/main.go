package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/SyadoWCS/single_tournament/database"
	"github.com/SyadoWCS/single_tournament/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// GORMでDB接続
	database.Connetct()

	// テンプレートファイル関連(HTML)
	list, err := template.New("t").ParseGlob("view/*.html")
	t := &Template{
		templates: template.Must(list, err),
	}

	// Echoのインスタンス作る
	e := echo.New()
	e.Renderer = t

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
