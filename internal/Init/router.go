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
	routes.SystemRouter(router)
	//用户路由
	routes.UserRouter(router)
	//文章路由

	_ = router.Run(constant.Config.System.Host + ":" + strconv.Itoa(constant.Config.System.Port))
	return router
}
