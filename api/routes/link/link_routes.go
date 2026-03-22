// Package link 友链模块路由（仅 GET / POST）
package link

import (
	linkcontroller "nectarpin/api/controllers/link"
	"nectarpin/api/middlewares"
	linkrepo "nectarpin/api/repositories/link"
	linkservice "nectarpin/api/services/link"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 注册友链相关路由
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	catRepo := linkrepo.NewFriendLinkCategoryRepository(db)
	linkRepo := linkrepo.NewFriendLinkRepository(db)
	pageRepo := linkrepo.NewFriendLinkPageRepository(db)
	svc := linkservice.NewFriendService(catRepo, linkRepo, pageRepo)
	ctrl := linkcontroller.NewFriendController(svc)

	publicAPI := router.Group("/api/public")
	{
		linkV1 := publicAPI.Group("/link/v1")
		{
			linkV1.GET("/page", middlewares.PublicHotResponseCache(), ctrl.GetPublicPage)
		}
	}

	protectedAPI := router.Group("/api/protected")
	protectedAPI.Use(middlewares.AuthMiddleware())
	{
		linkV1 := protectedAPI.Group("/link/v1")
		{
			linkV1.GET("/page/settings", ctrl.GetPageSettings)
			linkV1.POST("/page/settings", ctrl.SavePageSettings)

			linkV1.GET("/category/list", ctrl.ListCategories)
			linkV1.POST("/category/add", ctrl.CreateCategory)
			linkV1.POST("/category/update/:id", ctrl.UpdateCategory)
			linkV1.POST("/category/delete/:id", ctrl.DeleteCategory)

			linkV1.GET("/list", ctrl.ListFriendLinksAdmin)
			linkV1.POST("/add", ctrl.CreateFriendLink)
			linkV1.POST("/update/:id", ctrl.UpdateFriendLink)
			linkV1.POST("/delete/:id", ctrl.DeleteFriendLink)
		}
	}
}
