package controller

import (
	"gin/common"
	"gin/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取全部商品信息
func AllIdle(ctx *gin.Context) {
	DB := common.GetDB()

	var commodities []model.Commodity
	DB.Find(&commodities)
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(200, gin.H{
		"msg": commodities,
	})
}

// test0520 doSomething
func DoSomething(ctx *gin.Context) {
	DB := common.GetDB()

	var commodities []model.Commodity
	DB.Find(&commodities)
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(200, gin.H{
		"msg": commodities,
	})

}

//获取单个商品信息

// 发布闲置商品
func SellIdle(ctx *gin.Context) {
	DB := common.GetDB()

	// 从前端获取传来的数据
	telephone := ctx.PostForm("telephone")
	name := ctx.PostForm("name")
	picture1 := ctx.PostForm("picture1")
	price, _ := strconv.ParseFloat(ctx.PostForm("price"), 32)
	description := ctx.PostForm("description")

	//验证数据
	println(telephone)
	if len(telephone) != 11 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "手机号必填"})
		return
	}

	if len(name) == 0 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "请输入商品名称"})
		return
	}

	if len(picture1) == 0 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "请上传一张图片"})
		return
	}

	if price < 0 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "请输入合理的价格"})
		return
	}

	if len(description) == 0 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "请输入商品简介"})
		return
	}

	//查询商品是否已经存在
	if isIdleExist(DB, telephone, name) {
		ctx.JSON(422, gin.H{"code": 422, "msg": "该账户已存在相同商品"})
		return
	}

	newCommodity := model.Commodity{
		Telephone:   telephone,
		Name:        name,
		Picture1:    picture1,
		Price:       price,
		Description: description,
	}
	//在数据库中创建新的商品信息
	DB.Create(&newCommodity)
	ctx.JSON(200, gin.H{"code": 200, "msg": "成功发布闲置"})
}

// 使用telephone和商品名称name作为复合键查询商品是否已经存在，即同一个手机号不能有相同name的商品
func isIdleExist(db *gorm.DB, telephone string, name string) bool {
	var commodity model.Commodity
	db.Where("telephone = ? AND name = ?", telephone, name).First(&commodity)
	if commodity.ID != 0 {
		return true
	}
	return false
}
