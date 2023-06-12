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
	Id string `json:"id"`
}
type SendAddress struct {
	Id       string `json:"id"`
	Receiver string `json:"receiver"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
}
type SendGood struct {
	Id          string `json:"id"`
	User        string `json:"user"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Picture     string `json:"picture"`
	Price       string `json:"price"`
}
type Result struct {
	UserAddress SendAddress `json:"userAddress"`
	Goods       SendGood    `json:"goods"`
	Price       string      `json:"price"`
}

func GetFromCart(ctx *gin.Context) {

	var goodsInCart GoodsInCart
	if err := ctx.BindJSON(&goodsInCart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, is_Exist := ctx.Get("user")
	if is_Exist == false {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "user not exist"})
		return
	}
	userInfo := user.(model.User)

	DB := common.GetDB()

	//获取数据库的相关数据
	var good model.Goods
	e := DB.Table("goods").Where("id = ?", goodsInCart.Id).Find(&good)
	if e.Error != nil {
		fmt.Print(e.Error)
	}
	var address model.UserAddress
	e2 := DB.Table("user_addresses").Where("user_id = ?", userInfo.ID).Find(&address)
	if e2.Error != nil {
		fmt.Print(e.Error)
	}

	//填充发送的具体数据
	var sendGood SendGood
	sendGood.Id = strconv.Itoa(int(good.ID))
	sendGood.Name = good.Name
	sendGood.User = good.User
	sendGood.Description = good.Description
	sendGood.Picture = good.Picture
	sendGood.Price = good.Price
	//
	var sendAddress SendAddress
	sendAddress.Id = strconv.Itoa(int(address.ID))
	sendAddress.Receiver = address.Receiver
	sendAddress.Contact = address.Contact
	sendAddress.Address = address.Address
	//
	var result Result
	result.UserAddress = sendAddress
	result.Goods = sendGood
	result.Price = good.Price

	ctx.JSON(200, gin.H{
		"code":   "1",
		"msg":    "操作成功",
		"result": result,
	})

}
