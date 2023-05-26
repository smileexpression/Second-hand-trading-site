package controller

import (
	"gin/common"
	"gin/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AllIdle_2 struct { // "_2" 区分于commoditycontroller的AllIdle

	Id      string
	Name    string `gorm:"type:varchar(20);not null"`
	Picture string `gorm:"type:varchar(1024);not null"`
	Goods   []model.Goods
}

type SingleIdle struct {
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

func RecentIdle(ctx *gin.Context) {
	DB := common.GetDB()

	NUM := ctx.DefaultQuery("limit", "4")
	IntNum, err := strconv.Atoi(NUM) // 函数原型 ：func Atoi(s string) (int, error)
	if err != nil {
		print(err)
		//do not thing
	}
	//var recentGoods [4]model.Goods
	var recentGoods = make([]model.Goods, IntNum)
	var count int64
	DB.Table("goods").Count(&count)

	for i := int(count); i > int(count)-IntNum; i-- {
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

func ChooseCategory(ctx *gin.Context) {

	DB := common.GetDB()
	var result SingleIdle

	Cate_id := ctx.DefaultQuery("id", "3")

	var category model.Category
	DB.Table("categories").Where("id = ?", Cate_id).Find(&category)
	result.Id = category.Id
	result.Name = category.Name
	result.Picture = category.Picture

	var goods []model.Goods
	DB.Table("goods").Where("cate_Id = ?", Cate_id).Find(&goods)
	result.Goods = append(result.Goods, goods...)

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(200, gin.H{
		"code":   "1",
		"msg":    "获取分类下属物品成功",
		"result": result,
	})

}
