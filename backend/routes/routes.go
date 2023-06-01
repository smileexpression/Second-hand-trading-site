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

	home := r.Group("home")
	{
		home.GET("/goods", controller.GetGoods)

		home.GET("/banner", controller.GetBanner)
		home.GET("/new", controller.RecentIdle)
		home.POST("/order",controller.CreateOrder)
	}

	goods := r.Group("")
	{
		goods.GET("/goods", controller.GetOneGood)
	}
	return r
}
