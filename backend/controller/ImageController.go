package controller

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Arron2411616261/online/common"
	"github.com/Arron2411616261/online/model"
	"github.com/gin-gonic/gin"
)

// type Image struct {
// 	ID   uint `gorm:"primarykey"`
// 	Blob []byte
// }

func HandleUpload(c *gin.Context) {
	db := common.GetDB()
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var imageIds []uint
	for _, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {
			file, err := fileHeader.Open()
			if err != nil {
				log.Println("Error opening uploaded file:", err)
				continue
			}
			defer file.Close()
			blob, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println("Error reading uploaded file:", err)
				continue
			}
			dbImage := model.Image{Blob: blob}
			err = db.Create(&dbImage).Error
			if err != nil {
				log.Println("Error creating image record:", err)
				continue
			}
			imageIds = append(imageIds, dbImage.ID)
		}
	}
	c.JSON(http.StatusOK, gin.H{"imageIds": imageIds})
}

func HandleImage(c *gin.Context) {
	db := common.GetDB()
	id := c.Query("id")
	var image model.Image
	err := db.First(&image, id).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	c.Data(http.StatusOK, "image/jpeg", image.Blob)
}

func DeleteImage(c *gin.Context) {
	db := common.GetDB()
	id := c.Query("id")
	db.Where("id = ?", id).Delete(&model.Image{})
	c.JSON(http.StatusOK, gin.H{"message": "image deleted"})
}
