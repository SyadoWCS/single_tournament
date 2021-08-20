package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SyadoWCS/single_tournament/database"
	"github.com/SyadoWCS/single_tournament/model"
	"github.com/labstack/echo/v4"
)

type EntryInfo struct {
	TournamentId int
	Name         string
}

type TournamentEntry struct {
	TournamentId int
	Entry        []model.Entry
}

func EntryList(c echo.Context) error {
	var entry []model.Entry

	tournament_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// 大会IDを数字に変換できませんでした
		//return c.String(http.StatusBadRequest, "大会IDを数字に変換できませんでした")
	}

	// tournament_idに紐づく参加者一覧を取得
	database.DB.Where("tournament_id = ?", tournament_id).Find(&entry)

	/*if entry.TournamentId == 0 {
		// トーナメントIDを返すだけにしたい
	}*/

	tournament_entry := TournamentEntry{
		TournamentId: tournament_id,
		Entry:        entry,
	}

	json.Marshal(entry)
	return c.Render(http.StatusOK, "entry_list", tournament_entry)
}

func EntryNew(c echo.Context) error {
	tournament_id, err := strconv.Atoi(c.Param("tournament_id"))
	if err != nil {
		// 大会IDが取得できませんでした
	}
	fmt.Println("大会ID:" + c.Param("tournament_id"))

	entry_info := EntryInfo{
		TournamentId: tournament_id,
		Name:         "",
	}

	return c.Render(http.StatusOK, "entry_new", entry_info)
}

func EntryCreate(c echo.Context) error {
	name := c.FormValue("Name")
	tournament_id, err := strconv.Atoi(c.Param("tournament_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "大会IDを取得できませんでした")
	}

	entry := model.Entry{
		Name:         name,
		TournamentId: tournament_id,
	}
	// データ登録
	database.DB.Create(&entry)

	return c.Redirect(http.StatusFound, "/api/entry/list/"+c.Param("tournament_id"))
}

func EntryDelete(c echo.Context) error {
	tournament_id, err := strconv.Atoi(c.Param("tournament_id"))
	entry_id, err := strconv.Atoi(c.Param("entry_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "大会IDを取得できませんでした")
	}

	var entry model.Entry
	database.DB.Where("id = ? AND tournament_id = ?", entry_id, tournament_id).Delete(&entry)

	return c.Redirect(http.StatusFound, "/api/entry/list/"+c.Param("tournament_id"))
}
