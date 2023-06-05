package controller

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Image struct {
	ID   uint
	Blob []byte
}

type ImageController struct {
	db *gorm.DB
}

func NewImageController(db *gorm.DB) *ImageController {
	return &ImageController{db}
}

func (c *ImageController) UploadImage(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文件"})
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取文件"})
		return
	}

	image := Image{Blob: data}
	c.db.Create(&image)

	ctx.JSON(http.StatusOK, gin.H{"id": image.ID})
}

func (c *ImageController) GetImage(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "缺少参数"})
		return
	}

	var image Image
	result := c.db.First(&image, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到指定的图片"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
		}
		return
	}

	buffer := bytes.NewBuffer(image.Blob)

	ctx.Header("Content-Type", http.DetectContentType(image.Blob))
	ctx.Header("Content-Length", fmt.Sprintf("%d", len(image.Blob)))
	ctx.Stream(func(w io.Writer) bool {
		_, err := buffer.WriteTo(w)
		return err == nil
	})
}
