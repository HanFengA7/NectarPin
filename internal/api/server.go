package api

import (
	"context"
	"fmt"
	"log"
	"nectarpin/internal/config"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// Server HTTP 服务器
type Server struct {
	config     *config.Config
	router     *gin.Engine
	httpServer *http.Server
	startTime  time.Time
}

// NewServer 创建新的服务器实例
func NewServer(cfg *config.Config) *Server {
	gin.SetMode(cfg.Server.Mode)

	server := &Server{
		config:    cfg,
		router:    gin.New(),
		startTime: time.Now(),
	}

	server.setupMiddleware()
	server.setupRoutes()

	return server
}

// setupMiddleware 设置中间件
func (s *Server) setupMiddleware() {
	// 日志中间件
	s.router.Use(gin.Logger())
	// 恢复中间件
	s.router.Use(gin.Recovery())
}

// setupRoutes 设置路由
func (s *Server) setupRoutes() {
	// 健康检查
	s.router.GET("/health", func(c *gin.Context) {
		// 获取内存统计信息
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		memoryUsage := m.Alloc // 当前分配的内存（字节）

		c.JSON(http.StatusOK, gin.H{
			"status":     "ok",
			"message":    "NectarPin 服务运行正常",
			"version":    "1.0.0",
			"timestamp":  time.Now().Format(time.RFC3339),
			"uptime":     time.Since(s.startTime).String(),
			"memory":     fmt.Sprintf("%.2fMB", float64(memoryUsage)/1024/1024),
			"goroutines": runtime.NumGoroutine(),
			"database":   "connected",
		})
	})

	// API v1 路由组
	apiV1 := s.router.Group("/api/v1")
	{
		// 在这里添加你的路由
		// 例如: apiV1.GET("/moments", handlers.GetMoments)
		_ = apiV1 // 临时使用，防止未使用变量错误
	}
}

// Start 启动服务器
func (s *Server) Start() error {
	addr := s.config.Server.Host + ":" + s.config.Server.Port
	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("服务器启动在 %s", addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown 优雅关闭服务器
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
