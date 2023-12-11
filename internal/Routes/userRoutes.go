package Routes

import "github.com/gin-gonic/gin"

/*
UserRoutes

	登录接口
	注册接口
	用户信息查询接口
*/
func UserRoutes(router *gin.Engine) {
	user := router.Group("api/User")
	{
		user.GET("", func(c *gin.Context) {
			c.String(200, "hello-world")
		})
	}
}
