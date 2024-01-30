package Routes

import (
	"NectarPin/middleware"
	"github.com/gin-gonic/gin"
)

/*
ArticleRoutes [文章路由接口] [V1.0] [20240130]
------------------------------------------------------------------------------------------------------------------------

	╒[需要授权] [authArticleAPI]
	╞═todo [创建文章] [POST] [/api/Article/:id]
	╞═todo [编辑文章] [PUT] [/api/Article/:id]
	╞═todo [删除文章] [DELETE] [/api/Article/:id]
	╘═todo [扩展功能-文章访客分析] [GET] [/api/Article/func/analysis/:id]

------------------------------------------------------------------------------------------------------------------------

	╒[不需要授权] [notAuthUserAPI]
	╞═todo [查询文章-单篇] [GET] [/api/Article/:id]
	╞═todo [查询文章-单篇-加密] [GET] [/api/Article/:id/*password]
	╞═todo [查询文章-列表] [GET] [/api/Article/list/:pageSize/:pageNum]
	╞═todo [搜索文章] [POST] [/api/Article/search]
	╞═todo [扩展功能-文章点赞] [POST] [/api/Article/func/star]
	╘═todo [扩展功能-文章访客分析] [POST] [/api/Article/func/analysis]

------------------------------------------------------------------------------------------------------------------------
*/
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
		//todo [扩展功能-文章访客分析] [GET] [Private] [/api/Article/func/analysis/:id]
		authArticleAPI.GET("/func/analysis/:id")
	}

	//[不需要授权] [notAuthUserAPI]
	notAuthArticleAPI := router.Group("api/Article")
	{
		//todo [查询文章-单篇] [GET] [Public] [/api/Article/:id]
		notAuthArticleAPI.GET("/:id")
		//todo [查询文章-单篇-加密] [GET] [Public] [/api/Article/:id/*password]
		notAuthArticleAPI.GET("/:id/*password")
		//todo [查询文章-列表] [GET] [Public] [/api/Article/list/:pageSize/:pageNum]
		notAuthArticleAPI.GET("/list/:pageSize/:pageNum")
		//todo [搜索文章] [POST] [Public] [/api/Article/search]
		notAuthArticleAPI.POST("/search")
		//todo [扩展功能-文章点赞] [POST] [Public] [/api/Article/func/star]
		notAuthArticleAPI.POST("/func/star")
		//todo [扩展功能-文章访客分析] [POST] [Public] [/api/Article/func/analysis]
		notAuthArticleAPI.POST("/func/analysis")
	}
}
