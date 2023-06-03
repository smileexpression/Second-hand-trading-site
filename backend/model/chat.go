package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Me      string `gorm:"type:varchar(255);not null"`
	You     string `gorm:"type:varchar(255);not null"`
	Type    string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:varchar(255);not null"`
}
