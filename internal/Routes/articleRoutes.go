package Routes

import (
	"NectarPin/api"
	"NectarPin/middleware"
	"github.com/gin-gonic/gin"
)

/*
ArticleRoutes [文章路由接口] [V1.0] [20240130]
------------------------------------------------------------------------------------------------------------------------

	╒[需要授权] [authArticleAPI]
	╞═[1]═[创建文章] [POST] [/api/Article/]
	╞═[2]═[查询文章-列表-后台] [GET] [/api/Article/list/1/:pageSize/:pageNum]
	╞═[3]═todo [编辑文章] [PUT] [/api/Article/:id]
	╞═[4]═todo [删除文章] [DELETE] [/api/Article/:id]
	╘═[5]═todo [扩展功能-文章访客分析] [GET] [/api/Article/func/analysis/:id]

------------------------------------------------------------------------------------------------------------------------

	╒[不需要授权] [notAuthUserAPI]
	╞═[1]═todo [查询文章-单篇] [GET] [/api/Article/0/:id/:toType]
	╞═[2]═todo [查询文章-单篇-加密] [GET] [/api/Article/0/:id/:toType/*password]
	╞═[3]═todo [查询文章-列表-前台] [GET] [/api/Article/list/0/:pageSize/:pageNum]
	╞═[4]═todo [搜索文章] [POST] [/api/Article/search]
	╞═[5]═todo [扩展功能-文章点赞] [POST] [/api/Article/func/star]
	╘═[6]═todo [扩展功能-文章访客分析] [POST] [/api/Article/func/analysis]

------------------------------------------------------------------------------------------------------------------------
*/
func ArticleRoutes(router *gin.Engine) {

	// [需要授权] [authArticleAPI]
	authArticleAPI := router.Group("api/Article").Use(middleware.AuthJWT())
	{
		//[创建文章] [POST] [Private] [/api/Article]
		authArticleAPI.POST("", api.CreateArticle)
		//[查询文章-列表-后台] [GET] [Private] [/api/Article/list/1/:pageSize/:pageNum]
		authArticleAPI.GET("/list/1/:pageSize/:pageNum", api.GetArticleList1)
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
		//todo [查询文章-单篇] [GET] [Public] [/api/Article/0/:id/:toType]
		notAuthArticleAPI.GET("/0/:id/:toType", api.GetArticle0)
		//todo [查询文章-单篇-加密] [GET] [Public] [/api/Article/0/:id/:toType/*password]
		notAuthArticleAPI.GET("/0/:id/:toType/*password", api.GetArticle0)
		//todo [查询文章-列表-前台] [GET] [Public] [/api/Article/list/0/:pageSize/:pageNum]
		notAuthArticleAPI.GET("/list/0/:pageSize/:pageNum")
		//todo [搜索文章] [POST] [Public] [/api/Article/search]
		notAuthArticleAPI.POST("/search")
		//todo [扩展功能-文章点赞] [POST] [Public] [/api/Article/func/star]
		notAuthArticleAPI.POST("/func/star")
		//todo [扩展功能-文章访客分析] [POST] [Public] [/api/Article/func/analysis]
		notAuthArticleAPI.POST("/func/analysis")
	}
}
