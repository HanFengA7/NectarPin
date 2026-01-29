package routes

import "github.com/gin-gonic/gin"

// UserRoutes 用户相关路由
func UserRoutes(router *gin.Engine) {
	// 用户路由组[公共]
	userPublicGroup := router.Group("/api/public/user")
	{
		userPublicGroup.POST("/v1/register", nil)
		userPublicGroup.POST("/v1/login", nil)
	}

	// 用户路由组[私有]
	userPrivateGroup := router.Group("/api/private/user")
	{
		userPrivateGroup.GET("/v1/profile", nil)
	}

}
