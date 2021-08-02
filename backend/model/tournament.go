package model

import "gorm.io/gorm"

type Tournament struct {
	gorm.Model
	Name   string
	Number int
	Status int
}
