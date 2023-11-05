package Init

import (
	"NectarPin/constant"
	"NectarPin/internal/routes"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Router() *gin.Engine {
	router := gin.Default()

	//系统路由
	routes.SystemRoutes(router)
	//用户路由
	routes.UserRoutes(router)
	//文章路由
	//评论路由
	//友链路由
	//相册路由

	_ = router.Run(constant.Config.System.Host + ":" + strconv.Itoa(constant.Config.System.Port))
	return router
}
