package model

type Picture struct {
	GoodId   string
	Picture1 string `gorm:"type:varchar(1024);not null"`
	Picture2 string `gorm:"type:varchar(1024)"`
	Picture3 string `gorm:"type:varchar(1024)"`
	Picture4 string `gorm:"type:varchar(1024)"`
	Picture5 string `gorm:"type:varchar(1024)"`
}
