package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    // Echoのインスタンス作る
    e := echo.New()

    // ルーティング
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello World")
    })

    // サーバー起動
    e.Start(":80")
}