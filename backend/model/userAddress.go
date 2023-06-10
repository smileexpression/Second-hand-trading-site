package model

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model
	Receiver  string `json:"receiver" gorm:"type:varchar(255);not null"`
	Contact   string `json:"contact" gorm:"type:varchar(255);not null"`
	Address   string `json:"address" gorm:"type:varchar(1024);not null"`
	IsDefault bool   `json:"isDefault" gorm:"type:varchar(10);not null"`
}
