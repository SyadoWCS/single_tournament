package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Id           int    `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	Name         string `json:"name"`
	TournamentId int    `json:"tournament_id"`
	MatchId      int    `json:"match_id"`
}
