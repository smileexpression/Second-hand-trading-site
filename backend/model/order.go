package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model // ID gen update del
	Good_Id    string
	AddressId  string
	User_Id    uint
	PayMoney   int
}
