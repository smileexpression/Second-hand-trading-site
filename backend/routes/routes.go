package routes

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	category := r.Group("")
	{
		category.GET("/category", controller.ChooseCategory)
	}

	home := r.Group("home")
	{
		home.GET("/goods", controller.GetGoods)
		home.GET("/banner", controller.GetBanner)
		home.GET("/new", controller.RecentIdle)
	}
	return r
}
