package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// テスト用JSON
var userRegisterJSON = `{"first_name": "test", "last_name": "user", "email": "testuser@example.com", "password": "password", "password_confirm": "password"}`
var userLoginJSON = `{"email": "testuser@example.com", "password": "password"}`

func TestHome(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	Home(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello World", rec.Body.String())
}

func TestRegister(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(userRegisterJSON))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	Register(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestLogin(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(userLoginJSON))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	Login(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUser(t *testing.T) {

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/user", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	User(c)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.JSONEq(t, `{"message":"認証されていません"}`, rec.Body.String())
}

func TestLogout(t *testing.T) {

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/logout", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	Logout(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message":"ログアウト成功"}`, rec.Body.String())
}

/*func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}*/

func TestUserCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// 想定される値
	first_name := "test"
	last_name := "user"
	email := "testuser@example.com"
	password := "testuser"
	encode_password, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	// Mockの設定
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)").
		WithArgs(first_name, last_name, email, encode_password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// 実行
	res, err := UserCreate(first_name, last_name, email, password)
	if err != nil {
		t.Fatal(err)
	}

	if res.FirstName != first_name || res.LastName != last_name || res.Email != email {
		t.Errorf("取得結果不一致  %+v", res)
	}
}
