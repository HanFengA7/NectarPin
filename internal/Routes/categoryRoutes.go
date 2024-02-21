package Routes

import (
	"NectarPin/api"
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
	╞═[1]═[查询分类-单个] [GET] [/api/Category/:id]
	╘═[2]═[查询分类-列表] [GET] [/api/Category/list/:pageSize/:pageNum]

------------------------------------------------------------------------------------------------------------------------
*/
func CategoryRoutes(router *gin.Engine) {

	// [需要授权] [authArticleAPI]
	authCategoryAPI := router.Group("api/Category").Use(middleware.AuthJWT())
	{
		//[创建分类] [POST] [Private] [/api/Category]
		authCategoryAPI.POST("", api.CreateCategory)
		//[编辑分类] [PUT] [Private] [/api/Category/:id]
		authCategoryAPI.PUT("/:id", api.EditCategory)
		//[删除分类] [DELETE] [Private] [/api/Category/:id]
		authCategoryAPI.DELETE("/:id", api.DeleteCategory)
	}

	//[不需要授权] [notAuthUserAPI]
	notAuthCategoryAPI := router.Group("api/Category")
	{
		//[查询分类-单个] [GET] [Public] [/api/Category/:id]
		notAuthCategoryAPI.GET("/:id", api.GetCategory)
		//[查询分类-列表] [GET] [Public] [/api/Category/list/:pageSize/:pageNum]
		notAuthCategoryAPI.GET("/list/:pageSize/:pageNum", api.GetCategoryList)
	}
}
