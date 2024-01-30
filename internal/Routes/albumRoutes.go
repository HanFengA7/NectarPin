package Routes

import "github.com/gin-gonic/gin"

/*
AlbumRoutes [相册路由接口] [V1.0] []
------------------------------------------------------------------------------------------------------------------------

	╒[需要授权] [authAlbumAPI]
	╞═
	╞═
	╞═
	╘═

------------------------------------------------------------------------------------------------------------------------

	╒[不需要授权] [notAuthAlbumAPI]
	╞═
	╞═
	╞═
	╞═
	╞═
	╘═

------------------------------------------------------------------------------------------------------------------------
*/
func AlbumRoutes(router *gin.Engine) {
	album := router.Group("api/Album")
	{
		album.GET("", func(c *gin.Context) {
			c.String(200, "hello-world")
		})
	}
}
