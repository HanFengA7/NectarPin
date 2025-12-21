package api

import (
	"nectarpin/internal/api/routers"
)

// SetupRoutes 设置所有路由
func SetupRoutes(s *Server) {
	//系统路由
	routers.SystemRoutes(s.router, s.startTime)
	//用户路由
	routers.UserRoutes(s.router, s.Config)
}
