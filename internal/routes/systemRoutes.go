package routes

import (
	"NectarPin/api"
	"github.com/gin-gonic/gin"
)

func SystemRoutes(router *gin.Engine) {
	system := router.Group("api/System/")
	{
		system.GET("GetSystemInfo", api.GetSystemInfo)
	}
}
