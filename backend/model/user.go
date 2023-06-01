package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"nickname" gorm:"type:varchar(30);not null"`
	Telephone string `json:"account" gorm:"type:varchar(30);not null;unique"` //手机号就是账号，要求11位
	Password  string `json:"password" gorm:"size:255"`                        //密码长度：6-14
	Gender    string `json:"gender" gorm:"type:varchar(10);not null"`
	// Token     string `json:"token" gorm:"size:255;not null"`
	Avatar string `json:"avatar" gorm:"type:varchar(1024);not null"`
}
