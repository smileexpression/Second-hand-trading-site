package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model  //Id gen update del
	Cate_Id     string
	User        string `gorm:"type:varchar(30);not null"`
	Name        string `gorm:"type:varchar(50);not null"`
	Picture     string `gorm:"type:varchar(1024);not null"`
	Price       string `gorm:"type:float;not null"`
	Description string `gorm:"type:varchar(255);not null"`
	IsSold      *bool
}
