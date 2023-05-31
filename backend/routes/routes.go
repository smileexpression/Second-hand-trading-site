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
		apiCommodity.POST("/sellIdle", controller.SellIdle) //POST(½Ó¿Ú£¬º¯Êý),
		category.GET("/category", controller.ChooseCategory)
		apiCommodity.GET("/allIdle", controller.AllIdle)
	}

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.GET("/info", middleware.AuthMiddleware(), controller.Info)

	test_api := r.Group("api/ds")
	{
		test_api.GET("/doSomething", controller.DoSomething)
	}
	///home/goods

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
	return r
}
