package controller

import (
	"gin/common"
	"gin/model"

	"github.com/gin-gonic/gin"
)

func RecentIdle(ctx *gin.Context) {
	DB := common.GetDB()
	NUM := 4
	var recentGoods [4]model.Goods
	var count int64
	DB.Table("goods").Count(&count)

	for i := int(count); i > int(count)-NUM; i-- {
		print(int(count) - i)
		DB.Table("goods").Where("id = ?", i).Find(&recentGoods[int(count)-i])
	}
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(200, gin.H{
		"code":   "1",
		"msg":    "获取最近发布成功",
		"result": recentGoods,
	})

}
