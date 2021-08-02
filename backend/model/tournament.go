package model

import "gorm.io/gorm"

type Tournament struct {
	gorm.Model
	Id                int    `gorm:"primary_key"`
	Name              string `gorm:"size:255"`
	ParticipateNumber int
	Status            int `gorm:"not null"`
}
