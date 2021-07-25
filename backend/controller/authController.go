package controller

import (
	"net/http"

	"github.com/SyadoWCS/single_tournament/database"
	"github.com/SyadoWCS/single_tournament/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func Register(c echo.Context) error {
	var data map[string]string

	// リクエストデータをパースする
	if err := c.Bind(&data); err != nil {
		return err
	}

	// パスワードチェック
	if data["password"] != data["password_confirm"] {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "パスワードが一致しません",
		})
	}

	// パスワードをエンコード
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := model.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	// データ登録
	database.DB.Create(&user)

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	var data map[string]string
	var user model.User

	// リクエストデータをパースする
	if err := c.Bind(&data); err != nil {
		return err
	}

	// emailに紐づくユーザーを取得
	database.DB.Where("email = ?", data["email"]).First(&user)

	// データが存在しなかった場合は404エラーを返す
	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "ユーザが見つかりません",
		})
	}

	// パスワードが一致するかチェック
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "パスワードが間違っています",
		})
	}

	return c.JSON(http.StatusOK, user)
}
