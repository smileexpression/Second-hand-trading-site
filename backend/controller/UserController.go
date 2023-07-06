package controller

import (
	"fmt"
	"gin/common"
	"gin/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type apiAddress struct {
	AddressID string `json:"id"`
	Receiver  string `json:"receiver"`
	Contact   string `json:"contact"`
	Address   string `json:"address"`
}

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
		Avatar:    "1",
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
			"id":            newUser.ID,
			"account":       newUser.Telephone,
			"token":         token,
			"avatar":        newUser.Avatar,
			"nickname":      newUser.Name,
			"gender":        newUser.Gender,
			"userAddresses": nil,
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

	//获取user对应的所有地址
	var addressArray []model.UserAddress
	DB.Model(&model.UserAddress{}).Where("user_id=?", user.ID).Find(&addressArray)

	//for _, v := range addressArray {
	//	println("id", v.ID)
	//	println("Receiver", v.Receiver)
	//	println("Contact", v.Contact)
	//	println("Address", v.Address)
	//	println("UserID", v.UserID)
	//
	//}

	//如果user地址数组为空，则返回nil
	if len(addressArray) == 0 {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "登录成功",
			"result": gin.H{
				"id":            user.ID,
				"account":       user.Telephone,
				"token":         token,
				"avatar":        user.Avatar,
				"nickname":      user.Name,
				"gender":        user.Gender,
				"userAddresses": nil,
			},
		})
	} else {
		//创建响应数组
		var apiAddresses []apiAddress
		for _, address := range addressArray {
			resAddress := apiAddress{AddressID: strconv.Itoa(int(address.ID)), Receiver: address.Receiver, Contact: address.Contact, Address: address.Address}
			apiAddresses = append(apiAddresses, resAddress)
		}
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "登录成功",
			"result": gin.H{
				"id":            user.ID,
				"account":       user.Telephone,
				"token":         token,
				"avatar":        user.Avatar,
				"nickname":      user.Name,
				"gender":        user.Gender,
				"userAddresses": apiAddresses,
			},
		},
		)
	}
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
	DB := common.GetDB()
	pictureID, isSuccess := ctx.GetQuery("pictureID")
	user, is_Exist := ctx.Get("user")
	if !is_Exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "user not exist"})
		return
	}
	userInfo := user.(model.User)
	if !isSuccess {
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
	if userInfo.Avatar != pictureID {
		userInfo.Avatar = pictureID
	}
	DB.Model(&userInfo).Where("id=?", userInfo.ID).Update("avatar", userInfo.Avatar)
	ctx.JSON(200, gin.H{
		"code":     200,
		"msg":      "更换头像成功",
		"avatarID": userInfo.Avatar,
	})
}

type onPassword struct {
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}

func ChangePassword(ctx *gin.Context) {
	DB := common.GetDB()
	var receivePassword onPassword
	if err := ctx.BindJSON(&receivePassword); err != nil {
		ctx.JSON(422, gin.H{"code": 422, "msg": "获取失败"})
		return
	}
	user, is_Exist := ctx.Get("user")
	if !is_Exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "user not exist"})
		return
	}
	userInfo := user.(model.User)

	//验证新旧密码长度
	if len(receivePassword.Oldpassword) < 6 || len(receivePassword.Oldpassword) > 14 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "旧密码长度为6-14个字符"})
		return
	}

	if len(receivePassword.Newpassword) < 6 || len(receivePassword.Newpassword) > 14 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "新密码长度为6-14个字符"})
		return
	}

	//验证旧密码输入是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(receivePassword.Oldpassword)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "旧密码错误",
		})
		return
	}

	//如果旧密码输入正确且新密码长度符合，对新密码进行加密
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(receivePassword.Newpassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": "加密错误"})
		return
	}

	//更换密码
	//println("oldpassword:", userInfo.Password)
	userInfo.Password = string(HashPassword)
	//println("newpassword:", userInfo.Password)
	DB.Model(&userInfo).Where("id=?", userInfo.ID).Update("password", userInfo.Password)
	//返回响应
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "更换密码成功",
	})
}

func ChangeInfo(ctx *gin.Context) {
	DB := common.GetDB()
	var receiveInfo model.User
	if err := ctx.BindJSON(&receiveInfo); err != nil {
		ctx.JSON(422, gin.H{"code": 422, "msg": "获取失败"})
		return
	}

	newName := receiveInfo.Name
	newGender := receiveInfo.Gender

	if len(newName) == 0 || len(newGender) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "个人信息不完整，请重新填写",
		})
		return
	}
	user, is_Exist := ctx.Get("user")
	if !is_Exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "user not exist"})
		return
	}
	userInfo := user.(model.User)

	if userInfo.Name != newName || userInfo.Gender != newGender {
		userInfo.Name = newName
		userInfo.Gender = newGender
		DB.Model(&userInfo).Where("id=?", userInfo.ID).Updates(map[string]interface{}{"name": userInfo.Name, "gender": userInfo.Gender})
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "操作成功",
		"result": gin.H{
			"nickname": userInfo.Name,
			"gender":   userInfo.Gender,
		},
	})
}

func AddAddress(ctx *gin.Context) {
	DB := common.GetDB()
	var receiveAddress model.UserAddress
	if err := ctx.BindJSON(&receiveAddress); err != nil {
		ctx.JSON(422, gin.H{"code": 422, "msg": "获取失败"})
		return
	}

	if len(receiveAddress.Receiver) == 0 || len(receiveAddress.Contact) == 0 || len(receiveAddress.Address) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "地址信息不完整，请重新填写",
		})
		return
	}

	user, is_Exist := ctx.Get("user")
	if !is_Exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "user not exist"})
		return
	}
	userInfo := user.(model.User)
	receiveAddress.UserID = userInfo.ID

	DB.Create(&receiveAddress)

	var addressArray []model.UserAddress
	DB.Model(&model.UserAddress{}).Where("user_id=?", userInfo.ID).Find(&addressArray)

	//for _, v := range addressArray {
	//	println("id", v.ID)
	//	println("Receiver", v.Receiver)
	//	println("Contact", v.Contact)
	//	println("Address", v.Address)
	//	println("UserID", v.UserID)
	//
	//}

	var apiAddresses []apiAddress
	for _, address := range addressArray {
		resAddress := apiAddress{AddressID: strconv.Itoa(int(address.ID)), Receiver: address.Receiver, Contact: address.Contact, Address: address.Address}
		apiAddresses = append(apiAddresses, resAddress)
	}
	ctx.JSON(200, gin.H{
		"code":          200,
		"msg":           "新增收货地址成功",
		"userAddresses": apiAddresses,
	},
	)
}

func DeleteAddress(ctx *gin.Context) {
	DB := common.GetDB()
	addressID, isExist := ctx.GetQuery("id")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "没有传入参数id",
		})
		return
	}
	if len(addressID) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "id为空",
		})
		return
	}

	user, is_Exist := ctx.Get("user")
	if !is_Exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 400, "msg": "user not exist"})
		return
	}
	userInfo := user.(model.User)

	if err := DB.Where("id = ?", addressID).First(&model.UserAddress{}).Error; err == gorm.ErrRecordNotFound {
		//记录为空
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "该记录不存在"})
		return
	}

	DB.Where("id=?", addressID).Delete(&model.UserAddress{})

	var addressArray []model.UserAddress
	DB.Model(&model.UserAddress{}).Where("user_id=?", userInfo.ID).Find(&addressArray)

	//for _, v := range addressArray {
	//	println("id", v.ID)
	//	println("Receiver", v.Receiver)
	//	println("Contact", v.Contact)
	//	println("Address", v.Address)
	//	println("UserID", v.UserID)
	//
	//}

	var apiAddresses []apiAddress
	for _, address := range addressArray {
		resAddress := apiAddress{AddressID: strconv.Itoa(int(address.ID)), Receiver: address.Receiver, Contact: address.Contact, Address: address.Address}
		apiAddresses = append(apiAddresses, resAddress)
	}
	ctx.JSON(200, gin.H{
		"code":          200,
		"msg":           "删除收货地址成功",
		"userAddresses": apiAddresses,
	},
	)
}
