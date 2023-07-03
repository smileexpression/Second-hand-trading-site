package controller

import (
	"gin/common"
	"gin/model"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Image struct {
// 	ID   uint `gorm:"primarykey"`
// 	Blob []byte
// }

func HandleUpload(c *gin.Context) {
	db := common.GetDB()

	// 使用c.MultipartForm()从上下文中检索多部分表单数据
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 初始化一个空的图像ID切片，以跟踪成功上传的图像的ID
	var imageIds []uint

	// 循环遍历多部分表单数据中的所有文件头
	for _, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {
			// 对于每个文件头，使用fileHeader.Open()打开文件
			file, err := fileHeader.Open()
			if err != nil {
				log.Println("Error opening uploaded file:", err)
				continue
			}
			defer file.Close()
			// 使用ioutil.ReadAll()读取文件内容
			blob, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println("Error reading uploaded file:", err)
				continue
			}

			// 创建一个新的model.Image结构体，将文件内容存储在Blob字段中
			dbImage := model.Image{Blob: blob}
			err = db.Create(&dbImage).Error
			if err != nil {
				log.Println("Error creating image record:", err)
				continue
			}

			// 将成功上传的图像ID添加到imageIds切片中
			imageIds = append(imageIds, dbImage.ID)
		}
	}
	c.JSON(http.StatusOK, gin.H{"imageIds": imageIds})
}

func HandleImage(c *gin.Context) {
	db := common.GetDB()

	// 从上下文中获取查询参数"id"
	id := c.Query("id")

	// 声明一个model.Image变量image，用于存储从数据库中检索到的图像
	var image model.Image

	// 使用db.First()从数据库中检索具有指定ID的图像记录
	err := db.First(&image, id).Error

	// 如果检索过程中出现错误，使用c.AbortWithStatusJSON()返回一个JSON响应，指示无法找到图像
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	// 如果成功检索到图像记录，使用c.Data()将图像数据以HTTP响应的形式返回给客户端。此处假设所有图像均为JPEG格式的二进制数据，因此将MIME类型设置为"image/jpeg"
	c.Data(http.StatusOK, "image/jpeg", image.Blob)
}

func DeleteImage(c *gin.Context) {
	db := common.GetDB()
	id := c.Query("id")
	db.Where("id = ?", id).Delete(&model.Image{})
	c.JSON(http.StatusOK, gin.H{"message": "image deleted"})
}
