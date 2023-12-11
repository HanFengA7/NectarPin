package Routes

import "github.com/gin-gonic/gin"

func AlbumRoutes(router *gin.Engine) {
	album := router.Group("api/Album")
	{
		album.GET("", func(c *gin.Context) {
			c.String(200, "hello-world")
		})
	}
}
