package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model //Id gen update del
	Good_Id    string
	User_Id    string
}
