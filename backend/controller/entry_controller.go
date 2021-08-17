package controller

import (
	"encoding/json"
	"net/http"

	"github.com/SyadoWCS/single_tournament/database"
	"github.com/SyadoWCS/single_tournament/model"
	"github.com/labstack/echo/v4"
)

/*type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// テンプレートを描画
	return t.templates.ExecuteTemplate(w, name, data)
}*/

func EntryList(c echo.Context) error {
	var entry []model.Entry

	database.DB.Find(&entry)
	json.Marshal(entry)

	return c.Render(http.StatusOK, "entry_list", entry)
}

func EntryNew(c echo.Context) error {
	var entry model.Entry

	return c.Render(http.StatusOK, "entry_new", entry)
}
