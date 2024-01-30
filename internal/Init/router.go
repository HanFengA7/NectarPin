package Init

import (
	"NectarPin/constant"
	"NectarPin/internal/Routes"
	"NectarPin/middleware"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// 停止路由模块
func stopServer(router *gin.Engine) {
	// 执行一些清理操作

	// 关闭服务器
	if err := router.Run(":" + strconv.Itoa(constant.Config.System.Port)); err != nil {
		constant.Log.Error("服务关闭失败: ", err)
	} else {
		constant.Log.Info("服务已关闭")
	}
}

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
	//扩展路由

	// 启动服务（阻塞）
	srv := &http.Server{
		Addr:    constant.Config.System.Host + ":" + strconv.Itoa(constant.Config.System.Port),
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			constant.Log.Error("服务启动失败: ", err)
		}
	}()
	constant.Log.Info("路由启动成功")
	logrus.Infoln("路由启动成功")

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	return router
}
