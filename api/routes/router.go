// Package routes 提供应用程序路由配置
// 负责统一管理和注册所有模块的路由
package routes

import (
	"nectarpin/api/routes/article"
	"nectarpin/api/routes/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 配置应用程序所有路由
// 参数:
//   - router: Gin 路由引擎
//   - db: 数据库连接实例
//
// 该函数统一注册所有模块的路由:
//   - 用户模块路由 (/api/public/user/v1, /api/protected/user/v1)
//   - 文章模块路由 (/api/public/article/v1, /api/protected/article/v1)
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	user.SetupRoutes(router, db)
	article.SetupRoutes(router, db)
}
