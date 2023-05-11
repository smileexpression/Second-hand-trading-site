package routes

import (
	"gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	apiCommodity := r.Group("api/commodity")
	{
		apiCommodity.POST("/sellIdle", controller.SellIdle)
		apiCommodity.GET("/allIdle", controller.AllIdle)
	}

	return r
}
