package routes

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	apiCommodity := r.Group("api/commodity")
	{
		apiCommodity.POST("/sellIdle", controller.SellIdle) //POST(接口，函数)
		apiCommodity.GET("/allIdle", controller.AllIdle)
	}
	test_api := r.Group("api/ds")
	{
		test_api.GET("/doSomething", controller.DoSomething)
	}
	///home/goods

	home := r.Group("home")
	{
		home.GET("/goods", controller.GetGoods)
		home.GET("onegood", controller.GetOneGood)
	}

	goods := r.Group("")
	{
		goods.GET("/goods", controller.GetOneGood)
	}
	return r
}
