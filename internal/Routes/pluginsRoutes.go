package Routes

import (
	"NectarPin/internal/PluginCore"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PluginsRoutes(router *gin.Engine) {

	//[需要授权] [authUserAPI]
	authPluginsAPI := router.Group("api/Plugins")
	{
		authPluginsAPI.GET("/System/List", func(c *gin.Context) {

			data := PluginCore.PluginsListData
			for i, _ := range data {
				fmt.Println(data[i].PluginName)
			}
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": "200",
					"data": data[0].PluginName,
					"msg":  "[NectarPin]: 查询成功！",
				},
			)

		})
	}

	//[不需要授权] [notAuthUserAPI]
	notPluginsAPI := router.Group("api/Plugins")
	{

		notPluginsAPI.GET("/:name", func(c *gin.Context) {
			data := PluginCore.PluginsListData
			name := c.Param("name")
			for i, _ := range data {
				fmt.Println(data[i].PluginName)
				if name == data[i].PluginName {
					c.JSON(
						http.StatusOK,
						gin.H{
							"code": "200",
							"data": data[i],
							"msg":  "[NectarPin]: 查询成功！",
						},
					)
				} else {
					c.JSON(
						http.StatusOK,
						gin.H{
							"code": "200",
							"data": "null",
							"msg":  "[NectarPin]: 查询成功！",
						},
					)
				}
			}
		})
	}
}
