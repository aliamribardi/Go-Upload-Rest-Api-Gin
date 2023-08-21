package Models

import (
  		"gorm.io/gorm"
)

type Upload struct {
	ID			int
	NameFile	string		`gorm:"required"`
	gorm.Model
}