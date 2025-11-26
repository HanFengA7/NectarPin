package routers

import (
	"fmt"
	"nectarpin/internal/database"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// SystemRoutes 系统路由
func SystemRoutes(router *gin.Engine, startTime time.Time) {
	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		// 获取内存统计信息
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		memoryUsage := m.Alloc // 当前分配的内存（字节）

		// 检查数据库连接状态
		dbStatus := "disconnected"
		if database.DB != nil {
			sqlDB, err := database.DB.DB()
			if err == nil {
				if err := sqlDB.Ping(); err == nil {
					dbStatus = "connected"
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status":     "ok",
			"message":    "NectarPin 服务运行正常",
			"version":    "1.0.0",
			"timestamp":  time.Now().Format(time.RFC3339),
			"uptime":     time.Since(startTime).String(),
			"memory":     fmt.Sprintf("%.2fMB", float64(memoryUsage)/1024/1024),
			"goroutines": runtime.NumGoroutine(),
			"database":   dbStatus,
		})
	})
}
