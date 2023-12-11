package Routes

import "github.com/gin-gonic/gin"

func ArticleRoutes(router *gin.Engine) {
	article := router.Group("api/Article")
	{
		article.GET("", func(c *gin.Context) {
			c.String(200, "hello-world")
		})
	}
}
