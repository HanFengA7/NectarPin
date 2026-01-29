package bootstrap

import (
	"strconv"

	"NectarPin/api/routes"

	"github.com/gin-gonic/gin"
)

func InitRouter(port int) (*gin.Engine, error) {
	router := gin.Default()

	// 系统路由
	routes.SystemRoutes(router)
	// 用户路由
	routes.UserRoutes(router)

	if err := router.Run(":" + strconv.Itoa(port)); err != nil {
		return nil, err
	}
	return router, nil
}
