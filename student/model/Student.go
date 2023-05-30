package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name   string
	Age    uint8
	Gender uint8
}
