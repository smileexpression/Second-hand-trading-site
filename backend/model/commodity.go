package model

import "gorm.io/gorm"

type Commodity struct {
	gorm.Model
	Telephone   string  `gorm:"type:varchar(20);not null"` //用电话作为外键查询
	Name        string  `gorm:"type:varchar(50);not null"`
	Picture1    string  `gorm:"type:varchar(255);not null"`
	Price       float64 `gorm:"type:float;not null"`
	Description string  `gorm:"type:varchar(255);not null"`
}

//user实体类仅作为参考，可在mysql中看到这张表
