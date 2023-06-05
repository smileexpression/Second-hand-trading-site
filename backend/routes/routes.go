package routes

import (
	"gin/common"
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
	member := r.Group("member")
	{
		member.POST("/order", middleware.AuthMiddleware(), controller.CreateOrder)
		member.GET("/order/:id", middleware.AuthMiddleware(), controller.GetOrder)
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

	db := common.GetDB()
	imageController := controller.NewImageController(db)

	imageRoutes := r.Group("/images")
	{
		imageRoutes.POST("", imageController.UploadImage)
		imageRoutes.GET("/:id", imageController.GetImage)
	}

	return r
}
