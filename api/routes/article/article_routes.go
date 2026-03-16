// Package article 注册文章模块 GET/POST 路由
package article

import (
	articlecontroller "nectarpin/api/controllers/article"
	"nectarpin/api/middlewares"
	articlerepo "nectarpin/api/repositories/article"
	articleservice "nectarpin/api/services/article"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 注册文章相关路由
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	repo := articlerepo.NewArticleRepository(db)
	service := articleservice.NewArticleService(repo)
	ctrl := articlecontroller.NewArticleController(service)

	// 公开：列表、按 slug 获取详情、按 id 获取详情
	publicAPI := router.Group("/api/public")
	{
		articleV1 := publicAPI.Group("/article/v1")
		{
			// 列表
			articleV1.GET("/list", ctrl.List)
			// 按 slug 获取详情
			articleV1.GET("/infoBySlug/:slug", ctrl.GetBySlug)
			// 按 id 获取详情
			articleV1.GET("/infoById/:id", ctrl.GetByID)
		}
	}

	// 需认证：创建、编辑、删除
	protectedAPI := router.Group("/api/protected")
	protectedAPI.Use(middlewares.AuthMiddleware())
	{
		articleV1 := protectedAPI.Group("/article/v1")
		{
			// 创建
			articleV1.POST("/add", ctrl.Create)
			// 编辑
			articleV1.POST("/update/:id", ctrl.Update)
			// 删除
			articleV1.POST("/delete/:id", ctrl.Delete)
		}
	}
}
