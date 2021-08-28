package controller

import (
	"github.com/labstack/echo/v4"
)

// バックエンドのView配下のHTMLからPUTメソッドやDELETEメソッドを使用できるようにするため
func MethodOverride(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			method := c.Request().PostFormValue("_method")
			if method == "PUT" || method == "DELETE" {
				c.Request().Method = method
			}
		}
		return next(c)
	}
}
