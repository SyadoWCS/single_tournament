package model

import "gorm.io/gorm"

type Tournament struct {
	gorm.Model
	Id     int    `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	Name   string `json:"name"`
	People int    `json:"people"`
}
