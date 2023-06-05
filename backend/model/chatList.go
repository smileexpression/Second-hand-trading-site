package model

import "gorm.io/gorm"

type ChatList struct {
	gorm.Model
	Me  string `gorm:"type:varchar(255);not null"`
	You string `gorm:"type:varchar(255);not null"`
}
