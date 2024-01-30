package Routes

import (
	"NectarPin/api"
	"NectarPin/constant"
	"NectarPin/internal/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
SystemRoutes [系统路由接口] [V1.0] []
------------------------------------------------------------------------------------------------------------------------

	╒[需要授权] [authSystemAPI]
	╞═
	╞═
	╞═
	╘═

------------------------------------------------------------------------------------------------------------------------

	╒[不需要授权] [notAuthSystemAPI]
	╞═
	╞═
	╞═
	╞═
	╞═
	╘═

------------------------------------------------------------------------------------------------------------------------
*/
func SystemRoutes(router *gin.Engine) {
	system := router.Group("api/System/")
	{
		system.GET("GetSystemInfo", api.GetSystemInfo)

		system.GET("test", func(c *gin.Context) {
			var article []Models.Article
			err := constant.DB.Joins("User", constant.DB.Select("ID", "Username", "NickName", "Email", "AvatarUrl", "Role")).Preload("Category").Where("article.id = ?", 3).First(&article).Error
			if err != nil {
				panic("666")
			}
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"data":   article,
			})
		})
	}
}
