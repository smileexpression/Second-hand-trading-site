package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(30);not null"`
	Telephone string `gorm:"type:varchar(30);not null;unique"` //手机号就是账号，要求11位
	Password  string `gorm:"size:255"`                         //密码长度：6-14
}

//user实体类仅作为参考，可在mysql中看到这张表
