package controller

import (
	"github.com/Arron2411616261/online/common"

	"github.com/Arron2411616261/online/model"
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
