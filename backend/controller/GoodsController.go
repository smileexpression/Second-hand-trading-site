package controller

import (
	"gin/common"
	"gin/model"

	"github.com/gin-gonic/gin"
)

type AllIdle_2 struct { // "_2" 区分于commoditycontroller的AllIdle

	Id      string
	Name    string `gorm:"type:varchar(20);not null"`
	Picture string `gorm:"type:varchar(1024);not null"`
	Goods   []model.Goods
}

func GetGoods(ctx *gin.Context) {

	DB := common.GetDB()
	var result [4]AllIdle_2

	for i := 0; i < 4; i++ {

		var category model.Category
		DB.Table("categories").Where("id = ?", i+1).Find(&category)
		result[i].Id = category.Id
		result[i].Name = category.Name
		result[i].Picture = category.Picture

		var goods []model.Goods
		DB.Table("goods").Where("Cate_Id = ?", i+1).Find(&goods)
		result[i].Goods = append(result[i].Goods, goods...)
	}
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(200, gin.H{
		"code":   "1",
		"msg":    "获取全部商品成功",
		"result": result,
	})
}
