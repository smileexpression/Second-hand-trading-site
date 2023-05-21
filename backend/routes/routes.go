package routes

import (
	"gin/controller"
	"gin/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	apiCommodity := r.Group("api/commodity")
	{
		apiCommodity.POST("/sellIdle", controller.SellIdle)
		apiCommodity.GET("/allIdle", controller.AllIdle)
	}
	lr := r.Group("/LoginPage")
	{
		lr.POST("/login", controller.Login)
		lr.POST("/register", controller.Register)
		lr.GET("/info", middleware.AuthMiddleware(), controller.Info)
	}
	return r
}
