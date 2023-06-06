package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	//数据的规格参考goods表
	User_id string `gorm:"type:varchar(20);not null"`
	Good_id string `gorm:"type:varchar(20);not null"`
}
