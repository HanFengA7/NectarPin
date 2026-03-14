package bootstrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"nectarpin/api/models"
	"nectarpin/api/routes"
	"nectarpin/internal/utils"
)

// Application 应用程序主结构体
// 封装了应用运行所需的所有核心组件
type Application struct {
	Env      *Env        // 环境配置
	Gin      *gin.Engine // Gin 引擎实例
	Database *Database   // 数据库连接
}

// App 创建并初始化应用程序实例
// 返回:
//   - Application: 初始化完成的应用实例
//
// 初始化流程:
//  1. 加载环境配置
//  2. 建立数据库连接
//  3. 创建 Gin 引擎
//  4. 执行数据库迁移
//  5. 注册路由
func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Database = NewDatabase(app.Env)
	app.Gin = gin.Default()

	app.autoMigrate()
	routes.SetupRoutes(app.Gin, app.Database.DB)

	return *app
}

// Run 启动应用程序
// 启动 HTTP 服务器并监听关闭信号
//
// 该方法会:
//  1. 启动 Gin HTTP 服务器
//  2. 监听系统信号 (SIGINT, SIGTERM)
//  3. 收到关闭信号后优雅关闭
func (app *Application) Run() {
	port := app.Env.Config.Server.Port
	addr := fmt.Sprintf(":%d", port)

	utils.Logger.Infof("应用", "服务器正在启动，监听端口: %d", port)

	go func() {
		err := app.Gin.Run(addr)
		if err != nil {
			utils.Logger.Fatalf("应用", "服务器启动失败: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	utils.Logger.Info("应用", "正在关闭服务器...")
	app.Close()
	utils.Logger.Info("应用", "服务器已关闭")
}

// Close 关闭应用程序资源
// 释放数据库连接等资源
func (app *Application) Close() {
	if app.Database != nil {
		err := app.Database.Close()
		if err != nil {
			utils.Logger.Errorf("应用", "关闭数据库连接失败: %v", err)
		}
	}
}

// autoMigrate 执行数据库自动迁移
// 自动创建或更新数据库表结构
func (app *Application) autoMigrate() {
	err := app.Database.DB.AutoMigrate(
		&models.User{},
		&models.UserToken{},
	)
	if err != nil {
		utils.Logger.Fatalf("数据库", "自动迁移失败: %v", err)
	}
	utils.Logger.Success("数据库", "数据库表结构迁移完成")
}
