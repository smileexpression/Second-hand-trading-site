package model

type Banner struct {
	Id     string `gorm:"primaryKey"`
	ImgUrl string `gorm:"type:varchar(1024);not null"`
	HreUrl string `gorm:"type:varchar(1024);not null"`
}
