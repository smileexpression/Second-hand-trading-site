package controller

import (
	"crypto/rand"
	"fmt"
	"gin/common"
	"gin/model"
	"math/big"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type AllIdle struct { // "_2" 区分于commoditycontroller的AllIdle

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
	var result [4]AllIdle

	for i := 0; i < 4; i++ {

		var category model.Category
		DB.Table("categories").Where("id = ?", i+1).Find(&category)
		result[i].Id = category.Id
		result[i].Name = category.Name
		result[i].Picture = category.Picture

		var goods []model.Goods
		DB.Table("goods").Where("Cate_Id = ? AND is_sold=?", i+1, false).Find(&goods)
		result[i].Goods = append(result[i].Goods, goods...)
	}

	ctx.JSON(200, gin.H{
		"code":   "1",
		"msg":    "获取全部商品成功",
		"result": result,
	})
}

//暂且不考虑id转换错误

func GetOneGood(c *gin.Context) {
	db := common.GetDB()
	idStr := c.Query("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	var target model.Goods
	if err := db.Table("goods").Where("id = ?", id).First(&target).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"err warning": "Record of good no found",
			})
		}
	} else {
		var picTarget model.Picture
		if err := db.Table("pictures").Where("good_id = ?", target.ID).First(&picTarget).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{
					"err warning": "There is a error in pictures database,please contact the administrator",
				})
			}
		} else {
			//用户头像一般不会出错 简化代码不处理
			var user model.User
			db.Table("users").First(&user, target.User)
			p := [5]string{picTarget.Picture1, picTarget.Picture2, picTarget.Picture3, picTarget.Picture4, picTarget.Picture5}
			c.JSON(200, gin.H{
				"result":   target,
				"pictures": p,
				"user":     user,
			})
		}
	}

}

func RecentIdle(ctx *gin.Context) {
	DB := common.GetDB()

	NUM := ctx.DefaultQuery("limit", "4")
	IntNum, err := strconv.Atoi(NUM) // 函数原型 ：func Atoi(s string) (int, error)
	if err != nil {
		print(err)
		//do not thing
	}
	var count int64
	var ids []uint
	DB.Table("goods").Where("is_sold=?", false).Count(&count).Pluck("id", &ids)
	if IntNum > int(count) {
		IntNum = int(count) //让返回的数目不大于库存
	}
	var recentGoods = make([]model.Goods, IntNum)

	for i := int(count); i > int(count)-IntNum; i-- {
		print(int(count)-i, "fenge", i, "\n")
		DB.Table("goods").Where("id = ? AND is_sold=?", ids[i-1], false).Find(&recentGoods[int(count)-i])
	}

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
	DB.Table("goods").Where("cate_Id = ? AND is_sold=?", Cate_id, false).Find(&goods)
	result.Goods = append(result.Goods, goods...)

	ctx.JSON(200, gin.H{
		"code":   "1",
		"msg":    "获取分类下属物品成功",
		"result": result,
	})

}

type GoodInfo struct { //用于接收body参数
	Name        string   `json:"name"`
	Cate_id     string   `json:"cate_id"`
	Description string   `json:"description"`
	Picture     []string `json:"picture"`
	Price       string   `json:"price"`
}

func Release(ctx *gin.Context) {

	user, is_Exist := ctx.Get("user")
	if !is_Exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "user not exist"})
		return
	}
	userInfo := user.(model.User)

	//绑定body
	var goodInfo GoodInfo
	if err := ctx.BindJSON(&goodInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	DB := common.GetDB()

	//生成good
	userId := strconv.Itoa(int(userInfo.ID)) //之前将good表的User字段定义成了string
	good := model.Goods{
		Cate_Id:     goodInfo.Cate_id,
		User:        userId, //之前将good表的User字段定义成了string
		Name:        goodInfo.Name,
		Picture:     goodInfo.Picture[0],
		Price:       goodInfo.Price,
		Description: goodInfo.Description,
		Is_Sold:     false,
	}
	if err := DB.Create(&good).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	//生成good对应的picture
	goodId := strconv.Itoa(int(good.ID)) //之前将picture表goodId字段定义成了string
	picture := model.Picture{
		GoodId:   goodId,
		Picture1: goodInfo.Picture[0],
	}
	if err := DB.Select("good_id", "picture1").Create(&picture).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	//有点蠢的 当前picture表项添加多张图片
	pictureNum := len(goodInfo.Picture)
	switch pictureNum {
	case 2:
		DB.Model(&picture).Where("good_id=?", goodId).Select("picture2").Updates(map[string]interface{}{
			"picture2": goodInfo.Picture[1],
		})
	case 3:
		DB.Model(&picture).Where("good_id=?", goodId).Select("picture2", "picture3").Updates(map[string]interface{}{
			"picture2": goodInfo.Picture[1],
			"picture3": goodInfo.Picture[2],
		})
	case 4:
		DB.Model(&picture).Where("good_id=?", goodId).Select("picture2", "picture3", "picture4").Updates(map[string]interface{}{
			"picture2": goodInfo.Picture[1],
			"picture3": goodInfo.Picture[2],
			"picture4": goodInfo.Picture[3],
		})
	case 5:
		DB.Model(&picture).Where("good_id=?", goodId).Select("picture2", "picture3", "picture4", "picture5").Updates(map[string]interface{}{
			"picture2": goodInfo.Picture[1],
			"picture3": goodInfo.Picture[2],
			"picture4": goodInfo.Picture[3],
			"picture5": goodInfo.Picture[4],
		})

	}
	// if pictureNum > 1 {
	// 	DB.Model(&picture).Update("picture2", goodInfo.Picture[1])
	// 	if pictureNum > 2 {
	// 		DB.Model(&picture).Update("picture3", goodInfo.Picture[2])
	// 		if pictureNum > 3 {
	// 			DB.Model(&picture).Update("picture4", goodInfo.Picture[3])
	// 			if pictureNum > 4 {
	// 				DB.Model(&picture).Update("picture5", goodInfo.Picture[4])
	// 			}
	// 		}
	// 	}
	// }

	ctx.JSON(200, gin.H{
		"result": "Succeed",
	})
}

type apiGood struct {
	Id      uint   `json:"ID"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Price   string `json:"price"`
	Picture string `json:"picture"`
}

func transApiGood(good model.Goods) apiGood {
	return apiGood{Id: good.ID, Name: good.Name, Desc: good.Name, Price: good.Price, Picture: good.Picture}
}

func RecommendGoods(ctx *gin.Context) {
	DB := common.GetDB()
	str_limit := ctx.DefaultQuery("limit", "4")
	int_limit, err := strconv.Atoi(str_limit) // 函数原型 ：func Atoi(s string) (int, error)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get limit"})
		return
	}

	//未售出商品数量
	var isNotSold_Num int64
	DB.Table("goods").Where("is_sold=?", false).Count(&isNotSold_Num)
	//获得未售出的商品数组，数量为isNotSold_Num
	var notSoldGood_Array []model.Goods
	DB.Table("goods").Where("is_sold=?", false).Limit(int(isNotSold_Num)).Find(&notSoldGood_Array)
	println("isNotSold_Num", isNotSold_Num)

	//打印未售出的商品数组
	//println("notSoldGood_Array:")
	//for i, v := range notSoldGood_Array {
	//	println("k", i, ",v", v.ID)
	//}

	//通过比较获得最终展出商品的数量
	times := minNum(int_limit, int(isNotSold_Num))
	result := make([]apiGood, times)
	//记录已随机抽取商品的id
	var idRecord = make([]uint, times)
	//println("idRecord:")
	//for _, v := range idRecord {
	//	println(v)
	//}
	for i := 0; i < times; i++ {
		for {
			//随机抽取一个index,范围[0,total)
			ranIndex, _ := rand.Int(rand.Reader, big.NewInt(isNotSold_Num))
			//println("ranIndex=", ranIndex.Int64())
			index := uint(ranIndex.Int64())
			//str_ranID := strconv.Itoa(id)
			id := notSoldGood_Array[index].ID
			if checkRanID(idRecord, id) {
				idRecord[i] = id
				notSoldGood_Array[index].Is_Sold = true
				result[i] = transApiGood(notSoldGood_Array[index])
				break
			}
		}
	}

	//for i := range idRecord {
	//	println(idRecord[i])
	//}
	ctx.JSON(200, gin.H{
		"code":   200,
		"msg":    "操作成功",
		"result": result,
	})
}

func checkRanID(idRecord []uint, checkID uint) bool {
	//检查该id是否已经被抽取,如果已经被抽取，返回false
	for i := 0; i < len(idRecord); i++ {
		if idRecord[i] == checkID {
			return false
		}
	}
	return true
}

func minNum(limit int, total int) int {
	if limit <= total {
		return limit
	} else {
		return total
	}
}
