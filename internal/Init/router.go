package Init

import (
	"NectarPin/constant"
	"NectarPin/internal/Routes"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Router() *gin.Engine {
	router := gin.Default()

	//系统路由
	Routes.SystemRoutes(router)
	//用户路由
	Routes.UserRoutes(router)
	//文章路由
	Routes.ArticleRoutes(router)
	//评论路由
	Routes.CommentRoutes(router)
	//友链路由
	Routes.LinkRoutes(router)
	//相册路由
	Routes.AlbumRoutes(router)

	_ = router.Run(constant.Config.System.Host + ":" + strconv.Itoa(constant.Config.System.Port))
	return router
}
