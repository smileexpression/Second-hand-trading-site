package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 400,
					"data": nil,
					"msg":  err,
				})
			}
		}()

		ctx.Next()
	}
}
