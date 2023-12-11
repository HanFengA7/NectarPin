package Routes

import "github.com/gin-gonic/gin"

func CommentRoutes(router *gin.Engine) {
	comment := router.Group("api/Comment")
	{
		comment.GET("", func(c *gin.Context) {
			c.String(200, "hello-world")
		})
	}
}
