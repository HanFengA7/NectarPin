// Package article 注册文章模块路由（含分类、标签子路由）
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
	categoryRepo := articlerepo.NewArticleCategoryRepository(db)
	tagRepo := articlerepo.NewArticleTagRepository(db)

	svc := articleservice.NewArticleService(repo, categoryRepo, tagRepo)
	categorySvc := articleservice.NewArticleCategoryService(categoryRepo)
	tagSvc := articleservice.NewArticleTagService(tagRepo)

	ctrl := articlecontroller.NewArticleController(svc)
	categoryCtrl := articlecontroller.NewArticleCategoryController(categorySvc)
	tagCtrl := articlecontroller.NewArticleTagController(tagSvc)

	// ── 公开接口 ──
	publicAPI := router.Group("/api/public")
	{
		articleV1 := publicAPI.Group("/article/v1")
		{
			articleV1.GET("/list", ctrl.List)
			articleV1.GET("/infoBySlug/:slug", ctrl.GetBySlug)
			articleV1.GET("/infoById/:id", ctrl.GetByID)

			articleV1.GET("/category/list", categoryCtrl.List)
			articleV1.GET("/tag/list", tagCtrl.List)
		}
	}

	// ── 需认证接口 ──
	protectedAPI := router.Group("/api/protected")
	protectedAPI.Use(middlewares.AuthMiddleware())
	{
		articleV1 := protectedAPI.Group("/article/v1")
		{
			articleV1.GET("/list", ctrl.ListForAdmin)
			articleV1.GET("/infoById/:id", ctrl.GetByIDForAdmin)
			articleV1.POST("/add", ctrl.Create)
			articleV1.POST("/update/:id", ctrl.Update)
			articleV1.POST("/delete/:id", ctrl.Delete)

			articleV1.GET("/category/list", categoryCtrl.List)
			articleV1.POST("/category/add", categoryCtrl.Create)
			articleV1.POST("/category/update/:id", categoryCtrl.Update)
			articleV1.POST("/category/delete/:id", categoryCtrl.Delete)

			articleV1.GET("/tag/list", tagCtrl.List)
			articleV1.POST("/tag/add", tagCtrl.Create)
			articleV1.POST("/tag/update/:id", tagCtrl.Update)
			articleV1.POST("/tag/delete/:id", tagCtrl.Delete)
		}
	}
}
