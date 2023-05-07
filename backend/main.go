package main

import (
	"gin/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	r := gin.Default()

	//后端接口简单样例，具体路由设计可以看 https://blog.csdn.net/m0_51592186/article/details/118913401
	r.GET("/ping", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 前后端分离项目需要注意跨域问题，因此需要添加请求头
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
