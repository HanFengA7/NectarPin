package routes

import "github.com/gin-gonic/gin"

func UserRoutes(router *gin.Engine) {
	user := router.Group("api/User")
	{
		user.GET("", func(c *gin.Context) {
			c.String(200, "hello-world")
		})
	}
}
