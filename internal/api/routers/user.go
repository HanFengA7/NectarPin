package routers

import (
	"nectarpin/internal/api/handlers/user"
	"nectarpin/internal/api/middleware"
	"nectarpin/internal/config"

	"github.com/gin-gonic/gin"
)

// UserRoutes 用户路由
func UserRoutes(router *gin.Engine, cfg *config.Config) {
	// 公开路由（不需要认证）
	public := router.Group("/api/user")
	{
		public.POST("/register", user.UserRegister(cfg)) // 用户注册
		public.POST("/login", user.UserLogin(cfg))       // 用户登录
	}

	// 受保护路由（需要 JWT 认证）
	auth := router.Group("/api/user")
	auth.Use(middleware.JWTAuth(cfg))
	{
		auth.GET("/current", user.GetCurrentUser(cfg)) // 获取当前用户信息
		// 这里添加其他需要认证的用户路由
		// 例如: auth.PUT("/profile", user.UpdateProfile(cfg))
		// 例如: auth.PUT("/password", user.ChangePassword(cfg))
	}
}
