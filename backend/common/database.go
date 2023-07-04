package common

import (
	"fmt"
	"gin/model"
	"net/url"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//可以用navicat或datagrip等数据库操作软件，利用下面的信息登录数据库查看数据

	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to Db connection, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{}) // 此处创建了model文件夹下的user实体类，仅作参考
	db.AutoMigrate(&model.Goods{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Banner{})
	db.AutoMigrate(&model.Picture{})
	db.AutoMigrate(&model.Chat{})
	db.AutoMigrate(&model.ChatList{})
	db.AutoMigrate(&model.Cart{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.Image{})
	db.AutoMigrate(&model.UserAddress{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
