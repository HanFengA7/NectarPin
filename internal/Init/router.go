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
	routes.ArticleRoutes(router)
	//评论路由
	routes.CommentRoutes(router)
	//友链路由
	routes.LinkRoutes(router)
	//相册路由
	routes.AlbumRoutes(router)

	_ = router.Run(constant.Config.System.Host + ":" + strconv.Itoa(constant.Config.System.Port))
	return router
}
