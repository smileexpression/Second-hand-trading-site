package controller

import (
	"gin/common"
	"gin/model"

	"github.com/gin-gonic/gin"
)

func GetBanner(ctx *gin.Context) {

	DB := common.GetDB()
	var banners []model.Banner
	DB.Find(&banners)

	ctx.JSON(200, gin.H{
		"code":   "1",
		"msg":    "获取轮播图数据成功",
		"result": banners,
	})

}
