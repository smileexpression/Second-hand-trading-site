package model

type Address struct {
	Id           string `gorm:"primaryKey"`
	User_Id      uint
	Receiver     string
	Contact      string
	Address      string
	IsDefault    int
	FullLocation string
}
