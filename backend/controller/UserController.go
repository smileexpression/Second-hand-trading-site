package controller

import (
	"fmt"
	"gin/common"
	"gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 注册接口函数
func Register(ctx *gin.Context) {
	//获取数据
	DB := common.GetDB()
	var receiveUser model.User
	if err := ctx.BindJSON(&receiveUser); err != nil {
		ctx.JSON(422, gin.H{"code": 422, "msg": "获取失败"})
		return
	}
	//账号密码基本数据长度验证
	if len(receiveUser.Telephone) != 11 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(receiveUser.Password) < 6 || len(receiveUser.Password) > 14 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "密码长度需要设置为6-14个字符"})
		return
	}

	//验证手机号是否被注册过
	if isTelephoneExist(DB, receiveUser.Telephone) {
		ctx.JSON(422, gin.H{"code": 422, "msg": "该手机号已被注册"})
		return
	}
	//对密码进行加密处理
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(receiveUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	// 如果名称为空，随机创建一个名称
	// if len(receiveUser.Name) == 0 {
	// 	receiveUser.Name = RandomName(10)
	// }

	newUser := model.User{
		Gender: receiveUser.Gender,
		Name:   receiveUser.Name,
		// Token:     receiveUser.Token,
		Telephone: receiveUser.Telephone,
		Password:  string(HashPassword),
		Avatar:    "https://ts4.cn.mm.bing.net/th?id=OIP-C.Gve_dIeGxTpoiaU8CmdwOwHaHa&w=250&h=250&c=8&rs=1&qlt=90&o=6&dpr=1.5&pid=3.1&rm=2",
	}
	DB.Create(&newUser)
	//发放Token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"result": gin.H{
			"id":       newUser.ID,
			"account":  newUser.Telephone,
			"token":    token,
			"avatar":   newUser.Avatar,
			"nickname": newUser.Name,
			"gender":   newUser.Gender,
		},
	})
}

// 登录接口函数
func Login(ctx *gin.Context) {
	//获取参数
	DB := common.GetDB()
	var receiveUser model.User
	if err := ctx.BindJSON(&receiveUser); err != nil {
		ctx.JSON(422, gin.H{"code": 422, "msg": "获取失败"})
		return
	}

	//数据验证
	if len(receiveUser.Telephone) != 11 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "手机号为11位"})
		return
	}
	if len(receiveUser.Password) < 6 || len(receiveUser.Password) > 14 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "密码长度为6-14个字符"})
		return
	}

	//验证手机号对应的用户是否存在
	var user model.User
	DB.Where("telephone = ?", receiveUser.Telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(receiveUser.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}

	//发放token
	var err error
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"result": gin.H{
			"id":       user.ID,
			"account":  user.Telephone,
			"token":    token,
			"avatar":   user.Avatar,
			"nickname": user.Name,
			"gendar":   user.Gender,
		},
	})
}

// 验证手机号是否已被注册
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}

// // 随机创建用户名称
// func RandomName(n int) string {
// 	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
// 	result := make([]byte, n)
// 	rand.Seed(time.Now().Unix())
// 	for i := range result {
// 		result[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(result)
// }

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userinfo := user.(model.User)
	id := userinfo.ID
	fmt.Println(id)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}

func UpdateAvatar(ctx *gin.Context) {
	pictureID, isSuccess := ctx.GetQuery("pictureID")
	//user, is_Exist := ctx.Get("user")
	//if !is_Exist {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "user not exist"})
	//	return
	//}
	//userInfo := user.(model.User)
	if isSuccess == false {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "获取头像失败",
		})
		return
	}
	if pictureID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "头像不能为空",
		})
		return
	}
	//if userInfo.Avatar!=pictureID{
	//	userInfo.Avatar=pictureID
	//}
	ctx.JSON(200, gin.H{
		"code":     200,
		"msg":      "获得头像成功",
		"avatarID": pictureID,
	})
}
