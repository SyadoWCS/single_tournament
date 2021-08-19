package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/SyadoWCS/single_tournament/database"
	"github.com/SyadoWCS/single_tournament/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.StandardClaims
}

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

	var user model.User
	user, err := UserCreate(data["first_name"], data["last_name"], data["email"], data["password"])
	if err != nil {
		return err
	}

	/*user := model.User{
		FirstName: data["first_name"]
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	// データ登録
	database.DB.Create(&user)*/

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

	// JWT(Json Web Token)
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),            // stringに型変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// Cookie
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"jwt": token,
	})
}

func User(c echo.Context) error {
	// Loginで保存したJWTをCookieから取得する
	cookie, err := c.Cookie("jwt")
	if err != nil {
		// 401エラー
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "認証されていません",
		})
	}
	// token取得
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		// 401エラー
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "認証されていません",
		})
	}

	claims := token.Claims.(*Claims)
	// userIDを取得
	id := claims.Issuer

	var user model.User
	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(http.StatusOK, user)
}

func Logout(c echo.Context) error {
	// Cookie
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour) //マイナスにすると期限切れ扱い
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "ログアウト成功",
	})
}

func UserCreate(first_name string, last_name string, email string, not_encode_password string) (model.User, error) {
	// パスワードをエンコード
	password, _ := bcrypt.GenerateFromPassword([]byte(not_encode_password), 14)

	user := model.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  password,
	}

	// データ登録
	//database.DB.Create(&user)

	return user, database.DB.Create(&user).Error
}
