package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SyadoWCS/single_tournament/database"
	"github.com/SyadoWCS/single_tournament/model"
	"github.com/labstack/echo/v4"
)

func TournamentList(c echo.Context) error {
	var tournament []model.Tournament

	database.DB.Find(&tournament)
	json.Marshal(tournament)

	return c.Render(http.StatusOK, "tournament_list", tournament)
}

func TournamentNew(c echo.Context) error {
	var tournament model.Tournament

	return c.Render(http.StatusOK, "tournament_new", tournament)
}

func TournamentCreate(c echo.Context) error {
	name := c.FormValue("Name")
	people, err := strconv.Atoi(c.FormValue("People"))
	if err != nil {
		return c.String(http.StatusBadRequest, "大会を作成できませんでした")
	}

	tournament := model.Tournament{
		Name:   name,
		People: people,
	}
	// データ登録
	database.DB.Create(&tournament)

	return c.Redirect(http.StatusFound, "/tournament/list")
}

func TournamentEdit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "大会IDが見つかりませんでした")
	}

	tournament := model.Tournament{
		Id: id,
	}

	return c.Render(http.StatusOK, "tournament_edit", tournament)
}

func TournamentUpdate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "大会IDが見つかりませんでした")
	}
	name := c.FormValue("Name")
	people, err := strconv.Atoi(c.FormValue("People"))
	if err != nil {
		return c.String(http.StatusBadRequest, "参加人数が不正な値です")
	}

	tournament := model.Tournament{
		Name:   name,
		People: people,
	}

	database.DB.Model(model.Tournament{}).Where("id = ?", id).Updates(&tournament)

	return c.Redirect(http.StatusFound, "/tournament/list")
}

func TournamentDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "大会IDが見つかりませんでした")
	}

	var tournament model.Tournament
	var entry model.Entry
	database.DB.Where("id = ?", id).Delete(&tournament)
	database.DB.Where("tournament_id = ?", id).Delete(&entry)

	return c.Redirect(http.StatusFound, "/tournament/list")
}
