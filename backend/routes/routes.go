package routes

import (
	"gin/controller"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	category := r.Group("")
	{
		category.GET("/category", controller.ChooseCategory)
	}

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)

	home := r.Group("home")
	{
		home.GET("/goods", controller.GetGoods)

		home.GET("/banner", controller.GetBanner)
		home.GET("/new", controller.RecentIdle)
	}

	goods := r.Group("")
	{
		goods.GET("/goods", controller.GetOneGood)
	}

	chatList := r.Group("chat")
	{
		chatList.GET("/getmsg", middleware.AuthMiddleware(), controller.GetMsg)
		chatList.POST("/sendmsg", middleware.AuthMiddleware(), controller.SendMsg)
		chatList.POST("/addchat", middleware.AuthMiddleware(), controller.AddChat)
	}
	//member路由完善后可以将下面这个路由整合
	CartGroup := r.Group("member/cart")
	{
		CartGroup.POST("/add", middleware.AuthMiddleware(), controller.CartIn)
		CartGroup.GET("/pull", middleware.AuthMiddleware(), controller.CartOut)
		CartGroup.DELETE("/del", middleware.AuthMiddleware(), controller.CartDel)
	}
	return r
}
