package controller

import (
	"fmt"
	"gin/common"
	"gin/model"

	"github.com/gin-gonic/gin"
)

type SingleIdle struct {
	Id      string
	Name    string `gorm:"type:varchar(20);not null"`
	Picture string `gorm:"type:varchar(1024);not null"`
	Goods   []model.Goods
}

func ChooseCategory(ctx *gin.Context) {

	DB := common.GetDB()
	var result SingleIdle

	Cate_id := ctx.DefaultQuery("id", "3")

	var category model.Category
	DB.Table("categories").Where("id = ?", Cate_id).Find(&category)
	result.Id = category.Id
	result.Name = category.Name
	result.Picture = category.Picture
	fmt.Println(result.Name)

	var goods []model.Goods
	DB.Table("goods").Where("cate_Id = ?", Cate_id).Find(&goods)
	fmt.Println(goods)
	result.Goods = append(result.Goods, goods...)
	fmt.Println(result.Goods)

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(200, gin.H{
		"code":   "1",
		"msg":    "获取分类下属物品成功",
		"result": result,
	})

}
