package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Telephone string
	Password  string
}

//user实体类仅作为参考，可在mysql中看到这张表
