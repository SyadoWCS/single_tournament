package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

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
