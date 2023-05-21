package main

import (
	"gin/common"
	"gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = routes.CollectRoute(r)
	//后端接口简单样例，具体路由设计可以看 https://blog.csdn.net/m0_51592186/article/details/118913401

	r.Run(":8080")
	// 监听并在 0.0.0.0:8080 上启动服务
}
