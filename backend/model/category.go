package model

type Category struct {
	Id      string `gorm:"primaryKey"`
	Name    string `gorm:"type:varchar(50);not null"`
	Picture string `gorm:"type:varchar(1024);not null"`
}
