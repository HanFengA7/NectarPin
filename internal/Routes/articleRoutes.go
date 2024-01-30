package Routes

import (
	"NectarPin/middleware"
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(router *gin.Engine) {

	// [需要授权] [authArticleAPI]
	authArticleAPI := router.Group("api/Article").Use(middleware.AuthJWT())
	{
		//todo [创建文章] [POST] [Private] [/api/Article/:id]
		authArticleAPI.POST("/:id")
		//todo [编辑文章] [PUT] [Private] [/api/Article/:id]
		authArticleAPI.PUT("/:id")
		//todo [删除文章] [DELETE] [Private] [/api/Article/:id]
		authArticleAPI.DELETE("/:id")

	}

	//[不需要授权] [notAuthUserAPI]
	notAuthArticleAPI := router.Group("api/Article")
	{
		//todo [查询文章-单篇] [GET] [Public] [/api/Article/:id/:password]
		notAuthArticleAPI.GET("/:id/:password")
		//todo [查询文章-列表] [GET] [Public] [/api/Article/list/:pageSize/:pageNum]
		notAuthArticleAPI.GET("/list/:pageSize/:pageNum")
		//todo [搜索文章] [POST] [Public] [/api/Article/search]
		notAuthArticleAPI.POST("/search")
	}
}
