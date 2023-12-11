package Routes

import "github.com/gin-gonic/gin"

func LinkRoutes(router *gin.Engine) {
	link := router.Group("api/Link")
	{
		link.GET("", func(c *gin.Context) {
			c.String(200, "hello-world")
		})
	}
}
