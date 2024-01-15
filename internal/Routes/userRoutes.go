package Routes

import (
	"NectarPin/api"
	"github.com/gin-gonic/gin"
)

/*
UserRoutes

	登录接口
	注册接口
	用户信息查询接口
*/
func UserRoutes(router *gin.Engine) {
	user := router.Group("api/User")
	{
		user.GET("/add", api.CreateUser)
	}
}
