package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	schema = "%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local"
	// docker-compose.ymlに設定した環境変数を取得
	username       = os.Getenv("MYSQL_USER")
	password       = os.Getenv("MYSQL_PASSWORD")
	endpoint       = os.Getenv("MYSQL_ENDPOINT")
	dbName         = os.Getenv("MYSQL_DATABASE")
	datasourceName = fmt.Sprintf(schema, username, password, endpoint, dbName)
)

func main() {
	// GORMセット
	db, err := gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
	if err != nil {
		panic("DBに接続できません")
	}

	fmt.Println(db)
	fmt.Println("DBに接続できました")

	// Echoのインスタンス作る
	e := echo.New()

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	// サーバー起動
	e.Start(":80")
}
