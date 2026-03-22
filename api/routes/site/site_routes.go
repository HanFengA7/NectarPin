// Package site 站点相关路由（仅 GET / POST）
package site

import (
	sitecontroller "nectarpin/api/controllers/site"
	"nectarpin/api/middlewares"
	siterepo "nectarpin/api/repositories/site"
	siteservice "nectarpin/api/services/site"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 注册站点路由
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	repo := siterepo.NewSiteHomeRepository(db)
	svc := siteservice.NewSiteHomeService(repo)
	ctrl := sitecontroller.NewSiteHomeController(svc)

	publicAPI := router.Group("/api/public")
	{
		siteV1 := publicAPI.Group("/site/v1")
		{
			siteV1.GET("/home", ctrl.GetHome)
		}
	}

	protectedAPI := router.Group("/api/protected")
	protectedAPI.Use(middlewares.AuthMiddleware())
	{
		siteV1 := protectedAPI.Group("/site/v1")
		{
			siteV1.POST("/home", ctrl.SaveHome)
		}
	}
}
