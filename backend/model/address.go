package model

type Address struct {
	id           string `gorm:"primaryKey"`
	user_Id      uint
	receiver     string
	contact      string
	address      string
	isDefault    int
	fullLocation string
}
