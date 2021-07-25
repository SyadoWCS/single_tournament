package database

import (
	"fmt"
	"os"

	"github.com/SyadoWCS/single_tournament/model"
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
	// DBインスタンス
	DB *gorm.DB
)

func Connetct() {
	// GORMセット
	db, err := gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
	if err != nil {
		panic("DBに接続できません")
	}

	// コネクション情報追加(パッケージ外からDBにアクセスできるようにするため)
	DB = db

	db.AutoMigrate(&model.User{})
}
