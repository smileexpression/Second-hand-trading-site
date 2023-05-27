package controller

import (
	"fmt"
	"gin/common"
	"gin/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AllIdle_2 struct {
	Id      string
	Name    string `gorm:"type:varchar(20);not null"`
	Picture string `gorm:"type:varchar(1024);not null"`
	//Catagory model.Catagory
	Goods []model.Goods
}

func GetGoods(ctx *gin.Context) {

	DB := common.GetDB()
	var result [4]AllIdle_2
	var goods []model.Goods

	// i := 1
	// DB.First(&catagory, i)
	// result[i].Id = catagory.Id
	// result[i].Name = catagory.Name
	// result[i].Picture = catagory.Picture

	// i = 2
	// DB.First(&catagory, i)
	// result[i].Id = catagory.Id
	// result[i].Name = catagory.Name
	// result[i].Picture = catagory.Picture
	//var goods []model.Goods
	for i := 0; i < 4; i++ {
		var catagory model.Catagory
		DB.Table("catagories").Where("id = ?", i+1).Find(&catagory)

		// 	j := i + 1
		// 	DB.First(&catagory, j)
		result[i].Id = catagory.Id
		result[i].Name = catagory.Name
		result[i].Picture = catagory.Picture
		fmt.Println(catagory)

		DB.Table("goods").Where("Cata_Id = ?", i+1).Find(&goods)

		//result[i].Goods = goods
		//copy(result[i].Goods, goods)
		result[i].Goods = append(result[i].Goods, goods...)
		//fmt.Println(result[i].Goods)
	}
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(200, gin.H{
		"result": result,
	})
}

//暂且不考虑id转换错误

func GetOneGood(c *gin.Context) {
	db := common.GetDB()
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	//i := 1
	fmt.Println(1)
	if err != nil {
		fmt.Errorf("invalid id fomrat %v", err)
	}
	var target model.Goods
	db.Table("goods").Where("id = ?", id).First(&target)
	//if err != nil {
	//	fmt.Println("断点位于查表")
	//}
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"result": target,
	})
}
