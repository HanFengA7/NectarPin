package routers

import "github.com/gin-gonic/gin"

// UserRoutes 用户路由
func UserRoutes(router *gin.Engine) {
	// 用户登录
	router.POST("/user/login", func(ctx *gin.Context) {

	})
}
