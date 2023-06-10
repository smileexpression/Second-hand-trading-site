package controller

import (
	"fmt"
	"gin/common"
	"gin/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderInfo struct {
	Id        string `json:"goodId"` //goods_Id
	AddressId string `json:"addressId"`
}

func CreateOrder(ctx *gin.Context) {

	user, is_Exist := ctx.Get("user")
	if !is_Exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "user not exist"})
		return
	}
	userInfo := user.(model.User)

	var orderInfo OrderInfo
	//绑定结构体,接收body
	if err := ctx.BindJSON(&orderInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//生成订单
	var good model.Goods
	DB := common.GetDB()
	DB.Where("ID=?", orderInfo.Id).Find(&good)
	if good.Is_Sold {
		ctx.JSON(200, gin.H{
			"code": "0",
			"msg":  "手速太慢，商品被抢",
		})
		return
	}
	DB.Model(&good).Update("is_sold", true) //要不要考虑多线程？
	pMoney, _ := strconv.Atoi(good.Price)   //商品的价格为string，订单pay为int（字段类型）

	order := model.Order{
		Good_Id:   orderInfo.Id,
		AddressId: orderInfo.AddressId,
		User_Id:   userInfo.ID,
		PayMoney:  pMoney,
	}
	if err := DB.Create(&order).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
	//fmt.Print(good.Is_Sold)

	//返回id
	ctx.JSON(200, gin.H{
		"code": "1",
		"msg":  "操作成功",
		"id":   order.ID, //订单id
	})

}

func GetOrder(ctx *gin.Context) {

	orderId := ctx.Param("id")
	fmt.Print(orderId)
	DB := common.GetDB()

	var Order model.Order
	err := DB.Where("ID=?", orderId).Find(&Order)

	if err != nil { //应该没有
		// ctx.JSON(400, gin.H{
		// 	"code": "0",
		// 	"msg":  "订单不存在",
		// })
		fmt.Print(err)
	}
	ctx.JSON(200, gin.H{
		"code":      "1",
		"msg":       "操作成功",
		"countdown": "1800",
		"payMoney":  Order.PayMoney,
	})
	fmt.Print(Order.PayMoney)
	fmt.Print("???\n")
	fmt.Print(Order.AddressId)
}

type GoodsInCart struct {
	Id string `json:"goodId"`
}

func GetFromCart(ctx *gin.Context) {

	user, is_Exist := ctx.Get("user")
	if is_Exist == false {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "user not exist"})
		return
	}
	userInfo := user.(model.User)

	DB := common.GetDB()

	var goodsInCart GoodsInCart
	if err := ctx.BindJSON(&goodsInCart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

}
