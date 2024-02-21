package Routes

import (
	"NectarPin/middleware"
	"github.com/gin-gonic/gin"
)

/*
CategoryRoutes [文章路由接口] [V1.0] [20240221]
------------------------------------------------------------------------------------------------------------------------

	╒[需要授权] [authCategoryAPI]
	╞═[1]═[创建分类] [POST] [/api/Category]
	╞═[2]═[编辑分类] [PUT] [/api/Category/:id]
	╘═[3]═[删除分类] [DELETE] [/api/Category/:id]

------------------------------------------------------------------------------------------------------------------------

	╒[不需要授权] [notAuthCategoryAPI]
	╞═[1]═[查询分类-单个] [GET] [/api/Article/:id]
	╘═[2]═[查询分类-列表] [GET] [/api/Article/list/:pageSize/:pageNum]

------------------------------------------------------------------------------------------------------------------------
*/
func CategoryRoutes(router *gin.Engine) {

	// [需要授权] [authArticleAPI]
	authCategoryAPI := router.Group("api/Category").Use(middleware.AuthJWT())
	{
		//[创建分类] [POST] [Private] [/api/Category]
		authCategoryAPI.POST("")
		//[编辑分类] [PUT] [Private] [/api/Category/:id]
		authCategoryAPI.PUT("/:id")
		//[删除分类] [DELETE] [Private] [/api/Category/:id]
		authCategoryAPI.DELETE("/:id")
	}

	//[不需要授权] [notAuthUserAPI]
	notAuthCategoryAPI := router.Group("api/Article")
	{
		//[查询分类-单个] [GET] [Public] [/api/Article/:id]
		notAuthCategoryAPI.GET("/:id")
		//[查询分类-列表] [GET] [Public] [/api/Article/list/:pageSize/:pageNum]
		notAuthCategoryAPI.GET("/list/:pageSize/:pageNum")
	}
}
