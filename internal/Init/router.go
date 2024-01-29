package Init

import (
	"NectarPin/constant"
	"NectarPin/internal/Routes"
	"NectarPin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//[中间件][跨域]
	router.Use(middleware.Cors())

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
	//插件路由
	Routes.PluginsRoutes(router)
	go func() {
		_ = router.Run(constant.Config.System.Host + ":" + strconv.Itoa(constant.Config.System.Port))
	}()

	constant.Log.Info("路由启动成功")
	logrus.Infoln("路由启动成功")
	return router
}
