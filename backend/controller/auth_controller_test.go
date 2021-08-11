package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var userJSON = `{"first_name": "test", "last_name": "user", "email": "testuser@yahoo.co.jp", "password": "password", "password_confirm": "password"}`

func TestHome(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	Home(c)

	fmt.Println(rec)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello World", rec.Body.String())
}

func TestRegister(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(userJSON))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	Register(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

/*func TestLogin(t *testing.T) {

}*/

/*func TestUser(t *testing.T) {

}*/

/*func TestLogout(t *testing.T) {

}*/
