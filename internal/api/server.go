package api

import (
	"context"
	"log"
	"nectarpin/internal/config"
	"net/http"
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
	SetupRoutes(s)
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
