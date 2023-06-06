package controller

import (
	"fmt"
	"gin/common"
	"gin/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CartIn(c *gin.Context) {

	user, _ := c.Get("user")
	userinfo := user.(model.User)
	uId := userinfo.ID
	fmt.Println("uId: ", uId)
	db := common.GetDB()
	gId := c.Query("id")
	// 查询是否存在相同的数据
	var count int64
	db.Table("carts").Where("user_id = ? AND  good_id= ?", uId, gId).Count(&count)
	if count == 0 {
		// 不存在相同的数据，插入新数据
		var mId uint
		db.Table("carts").Select("COALESCE(MAX(id),0)").Scan(&mId)
		var re model.Cart
		re.ID = mId + 1
		re.User_id = strconv.Itoa(int(uId))
		re.Good_id = gId
		tx := db.Begin()

		if err := tx.Table("carts").Create(&re).Error; err != nil {
			// 处理错误
			tx.Rollback()
		}
		tx.Commit()
		fmt.Println("check 3")
		c.JSON(200, gin.H{
			"code":   "1",
			"result": "true",
			"msg":    "操作成功",
		})
	} else {
		// 存在相同的数据，不插入新数据
		c.JSON(200, gin.H{
			"code":   "0",
			"result": "flase",
			"msg":    "已经加入购物车",
		})
	}

}

func CartDel(c *gin.Context) {
	user, _ := c.Get("user")
	userinfo := user.(model.User)
	uId := userinfo.ID
	db := common.GetDB()
	gId := c.Query("id")

	var count int64
	db.Table("carts").Where("user_id = ? AND  good_id= ?", uId, gId).Count(&count)
	if count == 0 {
		c.JSON(200, gin.H{
			"code":   "0",
			"result": "flase",
			"msg":    "商品不在购物车内",
		})
	} else {
		if err := db.Table("carts").Where("user_id = ? AND  good_id= ?", uId, gId).Delete(&model.Cart{}).Error; err != nil {
			// 处理错误
		}
		//有必要可以先查询验证
		c.JSON(200, gin.H{
			"code":   "1",
			"result": "true",
			"msg":    "已删除",
		})
	}
}

func CartOut(c *gin.Context) {
	user, _ := c.Get("user")
	userinfo := user.(model.User)
	uId := userinfo.ID
	db := common.GetDB()

	var result []model.Goods
	err := db.Table("carts").Joins("left join goods ON carts.good_id = goods.id").Where("carts.user_id = ?", uId).Scan(&result)
	if err.Error != nil {
		//错误处理
	}
	if err.RowsAffected == 0 {
		//错误处理
	}
	// 输出查询结果
	c.JSON(200, gin.H{
		"code":   "1",
		"msg":    "获取购物车商品成功",
		"result": result,
	})
}
